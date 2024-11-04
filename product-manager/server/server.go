package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	hellopb "yorushika-store/product-manager/pkg/grpc"
	"yorushika-store/product-manager/repository"

	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
	sc  *ServerConfig
	l   net.Listener
	db  *repository.Database
}

type ServerConfig struct {
	Port string
}

func NewServer(sc *ServerConfig, db *repository.Database) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", sc.Port))
	if err != nil {
		log.Fatalf("net.Listen failed. err:%v", err)
	}

	srv := grpc.NewServer()
	hellopb.RegisterGreetingServiceServer(srv, NewMyServer(db))

	return &Server{
		srv: srv,
		sc:  sc,
		l:   listener,
		db:  db,
	}, nil
}

func NewServerConfig(port string) *ServerConfig {
	return &ServerConfig{
		Port: port,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		err := s.srv.Serve(s.l)
		if err != nil {
			log.Fatalf("server terminated %v", err)
		}
	}()
	log.Println("server started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	<-sig
	return nil
}

func (s *Server) Close() {
	s.srv.GracefulStop()
	log.Println("server stopped")
}

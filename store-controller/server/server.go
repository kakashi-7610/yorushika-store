package server

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	srv *http.Server
	sc  *ServerConfig
	l   net.Listener
}

type ServerConfig struct {
	Port   string
	Static string
}

type ProductServer struct {
	psc *ProductServerConfig
}

type ProductServerConfig struct {
	Host string
	Port string
}

func NewServer(sc *ServerConfig, psc *ProductServerConfig) (*Server, error) {
	mux, err := NewMux(sc, psc)
	if err != nil {
		log.Fatalf("NewMux failed, error: %v", err)
		return nil, err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", sc.Port))
	if err != nil {
		log.Fatalf("net.Listen failed. err:%v", err)
	}

	return &Server{
		srv: &http.Server{Handler: mux},
		sc:  sc,
		l:   listener,
	}, nil
}

func NewProductServer(psc *ProductServerConfig) *ProductServer {
	return &ProductServer{
		psc: psc,
	}
}

func NewServerConfig(port string, static string) *ServerConfig {
	return &ServerConfig{
		Port:   port,
		Static: static,
	}
}

func NewProductServerConfig(host string, port string) *ProductServerConfig {
	return &ProductServerConfig{
		Host: host,
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

func (s *Server) Close() error {
	err := s.srv.Close()
	if err != nil {
		return err
	}
	log.Println("server stopped")

	return nil
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

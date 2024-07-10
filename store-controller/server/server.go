package server

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	srv *http.Server
	sc  *ServerConfig
}

type ServerConfig struct {
	Hostname string
	Port     string
	Static   string
}

func NewServer(sc *ServerConfig) (*Server, error) {
	addr := fmt.Sprintf("%s:%s", sc.Hostname, sc.Port)
	mux, err := NewMux(sc)
	if err != nil {
		log.Fatalf("NewMux failed, error: %v", err)
		return nil, err
	}

	return &Server{
		srv: &http.Server{Addr: addr, Handler: mux},
		sc:  sc,
	}, nil
}

func NewServerConfig(hostname string, port string, static string) *ServerConfig {
	return &ServerConfig{
		Hostname: hostname,
		Port:     port,
		Static:   static,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		err := s.srv.ListenAndServe()
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

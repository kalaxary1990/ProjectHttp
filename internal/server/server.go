package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger  *log.Logger
	httpSrv *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.Handle("/upload", handlers.UploadHandler(logger))

	return &Server{
		logger: logger,
		httpSrv: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	s.logger.Printf("Starting server on %s", s.httpSrv.Addr)
	return s.httpSrv.ListenAndServe()
}

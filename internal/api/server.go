package api

import (
	"net/http"
	"note/internal/auth"
	"note/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() {
	logrus.Info("Running server")
	r := chi.NewRouter()

	noteService := service.NewNoteService()
	noteHandler := NewNoteHandler(noteService)
	noteHandler.RegisterRoutes(r)

	authService := auth.NewAuthService()
	authHandler := NewAuthHandler(authService)
	authHandler.RegisterAuth(r)

	http.ListenAndServe(s.addr, r)
}

package api

import (
	"database/sql"
	"net/http"
	"note/internal/auth"
	"note/internal/db"

	"note/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() {
	logrus.Info("Running server")
	r := chi.NewRouter()

	noteStore := db.NewNoteRepository(s.db)
	noteService := service.NewNoteService(noteStore)
	noteHandler := NewNoteHandler(noteService)
	noteHandler.RegisterRoutes(r)

	authService := auth.NewAuthService()
	authHandler := NewAuthHandler(authService)
	authHandler.RegisterAuth(r)

	err := http.ListenAndServe(s.addr, r)
	logrus.Fatal(err)
}

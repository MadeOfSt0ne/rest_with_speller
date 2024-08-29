package api

import (
	"net/http"
	"note/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	srv service.NoteService
}

func NewNoteHandler(srv service.NoteService) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {

		r.Get("/list", h.handleGetList)
		r.Post("/add", h.handleAddNote)
	})
}

func (h *Handler) handleAddNote(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Handling add note")
	answer := "add note successful"
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(answer))
}

func (h *Handler) handleGetList(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Handling get list")
	answer := "get list successful"
	w.Write([]byte(answer))
}

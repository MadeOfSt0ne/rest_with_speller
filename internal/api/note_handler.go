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
	r.Route("/api", func(router chi.Router) {
		router.Use(validateToken)
		router.Get("/list", h.handleGetList)
		router.Post("/add", h.handleAddNote)
	})
}

func (h *Handler) handleAddNote(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Handling add note")
	userId := req.Context().Value(AuthenticatedUserId).(int)
	logrus.Info("UserId from middleware:", userId)
	answer := "add note successful"
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(answer))
}

func (h *Handler) handleGetList(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Handling get list")
	answer := "get list successful"
	w.Write([]byte(answer))
}

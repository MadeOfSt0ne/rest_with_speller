package api

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"note/internal/service"
	"note/internal/types"

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

// Обработка запроса на добавление заметки
func (h *Handler) handleAddNote(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Handling add note")
	userId := req.Context().Value(AuthenticatedUserId).(int)
	var noteDto types.NoteDto
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		logrus.Error("error while reading from request body: ", err)
		http.Error(w, `{"error":"oops, something went wrong"}`, http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &noteDto); err != nil {
		logrus.Error("error while unmarshalling json: ", err)
		http.Error(w, `{"error":"oops, something went wrong"}`, http.StatusInternalServerError)
		return
	}

	note, err := h.srv.AddNewNote(userId, noteDto)
	if err != nil {
		logrus.Error("error while adding note: ", err)
		http.Error(w, `{"error":"oops, something went wrong"}`, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(note)
	if err != nil {
		logrus.Error("error while marshalling note: ", err)
		http.Error(w, `{"error":"oops, something went wrong"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application-json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(response))
	if err != nil {
		slog.Error("failed to write the response.", "err", err)
	}
}

// Обработка запроса на получение списка заметок
func (h *Handler) handleGetList(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Handling get list")
	userId := req.Context().Value(AuthenticatedUserId).(int)
	notes, err := h.srv.GetNotes(userId)
	if err != nil {
		logrus.Error("error while getting notes: ", err)
		http.Error(w, `{"error":"oops, something went wrong"}`, http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(notes)
	if err != nil {
		logrus.Error("failed to marshal tasks: ", err)
		http.Error(w, `{"error":"oops, something went wrong"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application-json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		slog.Error("failed to write the response.", "err", err)
	}
}

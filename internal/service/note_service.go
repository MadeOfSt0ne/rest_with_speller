package service

import (
	"fmt"
	"note/internal/types"

	"github.com/sirupsen/logrus"
)

type NoteService struct {
	store types.NoteStore
}

func NewNoteService(store types.NoteStore) NoteService {
	return NoteService{store: store}
}

// Добавление новой заметки в репозиторий
func (n NoteService) AddNewNote(userId int, note types.NoteDto) (types.Note, error) {
	var empty types.Note
	if note.Text == "" {
		logrus.Info("Empty text!")
		return empty, fmt.Errorf("text should not be empty")
	}

	editedText, err := useSpellChecker(note.Text)
	if err != nil {
		logrus.Info("speller returned error: ", err)
		return empty, fmt.Errorf("something went wrong")
	}
	note.Text = editedText

	addedNote, err := n.store.AddNewNote(userId, note)
	return addedNote, err
}

// Запрос заметок из репозитория
func (n NoteService) GetNotes(userId int) ([]types.Note, error) {
	notes, err := n.store.GetAllNotes(userId)
	if err != nil {
		logrus.Error("repository returned error")
		return notes, err
	}
	return notes, nil
}

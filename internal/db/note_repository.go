package db

import "note/internal/types"

type NoteRepository struct {
}

func NewNoteRepository() *NoteRepository {
	return &NoteRepository{}
}

var notes = map[int]types.Note{}
var id int

// Добавление заметки
func (n NoteRepository) AddNewNote(userId int, note types.NoteDto) (types.Note, error) {
	id++
	notes[id] = types.Note{Id: id, Author_id: userId, Title: note.Title, Text: note.Text}
	return notes[id], nil
}

// Получение списка заметок пользователя
func (n NoteRepository) GetAllNotes(userId int) ([]types.Note, error) {
	var noteList []types.Note
	for _, val := range notes {
		if val.Author_id == userId {
			noteList = append(noteList, val)
		}
	}
	return noteList, nil
}

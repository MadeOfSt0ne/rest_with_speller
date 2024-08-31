package db

import (
	"database/sql"
	"fmt"
	"note/internal/types"

	"github.com/sirupsen/logrus"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

// Добавление заметки
func (n NoteRepository) AddNewNote(userId int, note types.NoteDto) (types.Note, error) {
	logrus.Infof("adding new note: %s for user: %d", note, userId)
	var newNote types.Note
	res, err := n.db.Exec("INSERT INTO notes (author_id, title, text) VALUES (:author_id, :title, :text)",
		sql.Named("author_id", userId),
		sql.Named("title", note.Title),
		sql.Named("text", note.Text))
	if err != nil {
		logrus.Error("error while adding new note: ", err)
		return newNote, fmt.Errorf("something went wrong")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logrus.Error("error while getting id of new note: ", err)
		return newNote, fmt.Errorf("something went wrong")
	}
	newNote.Id = int(id)
	newNote.Author_id = userId
	newNote.Title = note.Title
	newNote.Text = note.Text

	return newNote, nil
}

// Получение списка заметок пользователя
func (n NoteRepository) GetAllNotes(userId int) ([]types.Note, error) {
	logrus.Info("getting note for user: ", userId)
	noteList := []types.Note{}
	rows, err := n.db.Query("SELECT id, author_id, title, text FROM notes WHERE author_id = :user_id ORDER BY id",
		sql.Named("user_id", userId))
	if err != nil {
		logrus.Error("error while getting notes: ", err)
		return nil, fmt.Errorf("something went wrong")
	}
	defer rows.Close()

	for rows.Next() {
		v := types.Note{}
		err := rows.Scan(&v.Id, &v.Author_id, &v.Title, &v.Text)
		if err != nil {
			logrus.Error("error while scanning rows: ", err)
			return nil, fmt.Errorf("something went wrong")
		}
		noteList = append(noteList, v)
	}
	if err := rows.Err(); err != nil {
		logrus.Error("error with rows: ", err)
		return nil, fmt.Errorf("something went wrong")
	}
	return noteList, nil
}

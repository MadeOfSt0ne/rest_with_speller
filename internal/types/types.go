package types

// Структура заметки
type Note struct {
	Id        int    `json:"id"`
	Author_id int    `json:"author_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}

// DTO заметки
type NoteDto struct {
	Title string
	Text  string
}

// Структура пользователя
type User struct {
	Id       int    `json:"Id"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

// Структура для логина
type LoginInfo struct {
	Login    string
	Password string
}

// Интерфейс репозитория
type NoteStore interface {
	GetAllNotes(userId int) ([]Note, error)
	AddNewNote(userId int, note NoteDto) (Note, error)
}

// Структура ответа от спеллера
type CheckText struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

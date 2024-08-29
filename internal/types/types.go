package types

// Структура заметки
type Note struct {
	Id        int    `json:"id"`
	Author_id int    `json:"author_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
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

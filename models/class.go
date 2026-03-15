package models
type Class struct {
	ID      int    `json:"id"`
	NAME    string `json:"name"`

	TeacherID int    `json:"teacher_id"`
}
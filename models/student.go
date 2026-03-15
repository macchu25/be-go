package models
type Student struct {
	ID      int    `json:"id"`
	NAME    string `json:"name"`
	EMAIL     string    `json:"email"`
	ClassID int    `json:"class_id"`
}
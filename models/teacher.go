package models
type Teacher struct {
	ID      int    `json:"id"`
	
	EMAIL     string    `json:"email"`
	
	NAME    string `json:"name"`
	PHONE string    `json:"phone"`
}
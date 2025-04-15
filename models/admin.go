package models

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Service     string `json:"service"`
	Message     string `json:"message"`
	SubmittedAt string `json:"submitted_at"`
	Status      string `json:"status"`
}

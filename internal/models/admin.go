package models

type Admin struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

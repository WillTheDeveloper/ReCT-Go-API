package main

type user struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Github  string `json:"github"`
	Email   string `json:"email"`
	Discord string `json:"discord"`
}

type project struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"private"`
	Language    string `json:"language"`
	Creator     string `json:"creator"`
}

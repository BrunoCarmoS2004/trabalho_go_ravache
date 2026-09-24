package models

type Aluno struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
	Email     string `json:"email"`
}
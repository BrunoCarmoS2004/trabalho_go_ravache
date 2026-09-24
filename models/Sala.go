package models

type Sala struct {
	Id         int      `json:"id"`
	Nome       string   `json:"nome"`
	Capacidade int      `json:"capacidade"`
	Recursos   []string `json:"recursos"`
}
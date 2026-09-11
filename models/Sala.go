package models

type Sala struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Capacidade string `json:"capacidade"`
	Recurso string `json:"recurso"`
}
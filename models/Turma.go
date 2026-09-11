package models

type Turma struct{
	Id int `json:"id"`
	Nome string `json:"nome"`
	Descricao string `json:"decricao"`
	QuantidadeAlunos int `json:"quantidadeAlunos"`
	Alunos [] Aluno `json:"alunos"`
	Sala Sala `json:"sala"`
}
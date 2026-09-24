package models

type Turma struct {
	Id              int     `json:"id"`
	Nome            string  `json:"nome"`
	Disciplina      string  `json:"disciplina"`
	Professor       string  `json:"professor"`
	QuantidadeAlunos int    `json:"quantidadeAlunos"`
	Alunos          []Aluno `json:"alunos"`
	SalaId          int     `json:"salaId"`
	DiaSemana       string  `json:"diaSemana"`
	HorarioInicio   string  `json:"horarioInicio"`
	HorarioFim      string  `json:"horarioFim"`
}
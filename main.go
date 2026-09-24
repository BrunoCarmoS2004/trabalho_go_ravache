package main

import (
	"net/http"
	"strconv"
	"time"

	"TrabalhoFinalGo/models"

	"github.com/gin-gonic/gin"
)

var alunos []models.Aluno
var salas []models.Sala
var turmas []models.Turma
var proximoIdAluno = 1
var proximoIdSala = 1
var proximoIdTurma = 1

func main() {

	r := gin.New()

	// Uso dos Middlewares globais nativos e personalizados
	r.Use(gin.Recovery())

	// 4. Mapeamento de Rotas sob Grupo Versionado
	v1 := r.Group("/api/v1")
	{
		// Monitoramento da API
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "healthy",
				"timestamp": time.Now(),
				"version":   "1.0.0",
			})
		})

		v1.POST("/alunos", criarAluno)
		v1.GET("/alunos", listarAlunos)
		v1.GET("/alunos/:id", buscarAlunoPorId)
		v1.PUT("/alunos/:id", atualizarAluno)
		v1.DELETE("/alunos/:id", deletarAluno)

		v1.POST("/salas", criarSala)
		v1.GET("/salas", listarSalas)
		v1.GET("/salas/:id", buscarSalaPorId)
		v1.PUT("/salas/:id", atualizarSala)
		v1.DELETE("/salas/:id", deletarSala)

		v1.POST("/turmas", criarTurma)
		v1.GET("/turmas", listarTurmas)
		v1.GET("/turmas/:id", buscarTurmaPorId)
		v1.PUT("/turmas/:id", atualizarTurma)
		v1.DELETE("/turmas/:id", deletarTurma)

		v1.POST("/turmas/:id/matricular", matricularAluno)
		v1.GET("/turmas/:id/alunos", listarAlunosDaTurma)
		v1.POST("/turmas/:id/alocar", alocarSala)
	}

	r.Run(":8080")
}

func criarSala(c *gin.Context) {
	var sala models.Sala

	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if sala.Capacidade <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capacidade deve ser maior que zero"})
		return
	}

	sala.Id = proximoIdSala
	proximoIdSala++

	salas = append(salas, sala)

	c.JSON(http.StatusCreated, sala)
}

func criarTurma(c *gin.Context) {
	var turma models.Turma

	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	turma.Id = proximoIdTurma
	proximoIdTurma++

	turmas = append(turmas, turma)

	c.JSON(http.StatusCreated, turma)
}

func criarAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	aluno.Id = proximoIdAluno
	proximoIdAluno++

	alunos = append(alunos, aluno)

	c.JSON(http.StatusCreated, aluno)
}

func listarAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, alunos)
}

func listarSalas(c *gin.Context) {
	c.JSON(http.StatusOK, salas)
}

func listarTurmas(c *gin.Context) {
	for indice := range turmas {
		turmas[indice].QuantidadeAlunos = len(turmas[indice].Alunos)
	}

	c.JSON(http.StatusOK, turmas)
}

func buscarAlunoPorId(c *gin.Context) {
	indice := buscarPorIdAluno(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, alunos[indice])
}

func buscarSalaPorId(c *gin.Context) {
	indice := buscarPorIdSala(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrado"})
		return
	}

	c.JSON(http.StatusOK, salas[indice])
}

func buscarTurmaPorId(c *gin.Context) {
	indice := buscarPorIdTurma(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrado"})
		return
	}

	c.JSON(http.StatusOK, turmas[indice])
}



func atualizarAluno(c *gin.Context) {
	indice := buscarPorIdAluno(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	aluno.Id = alunos[indice].Id
	alunos[indice] = aluno

	c.JSON(http.StatusOK, aluno)
}

func atualizarSala(c *gin.Context) {
	indice := buscarPorIdSala(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrado"})
		return
	}

	var sala models.Sala

	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	sala.Id = salas[indice].Id
	salas[indice] = sala

	c.JSON(http.StatusOK, sala)
}

func atualizarTurma(c *gin.Context) {
	indice := buscarPorIdTurma(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrado"})
		return
	}

	var turma models.Turma

	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	turma.Id = turmas[indice].Id
	turmas[indice] = turma

	c.JSON(http.StatusOK, turma)
}

func deletarAluno(c *gin.Context) {
	indice := buscarPorIdAluno(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	alunos = append(alunos[:indice], alunos[indice+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Aluno deletado"})
}

func deletarSala(c *gin.Context) {
	indice := buscarPorIdSala(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrado"})
		return
	}

	salas = append(salas[:indice], salas[indice+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Sala deletada"})
}

func deletarTurma(c *gin.Context) {
	indice := buscarPorIdTurma(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrado"})
		return
	}

	turmas = append(turmas[:indice], turmas[indice+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Turma deletada"})
}


func matricularAluno(c *gin.Context) {
	indiceTurma := buscarPorIdTurma(c)

	if indiceTurma == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrada"})
		return
	}

	var body struct {
		AlunoId int `json:"aluno_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if alunoNaoExiste(body.AlunoId) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	if alunoEstaNaTurma(turmas[indiceTurma], body.AlunoId) {
		c.JSON(http.StatusConflict, gin.H{"error": "Aluno já matriculado na turma"})
		return
	}

	if turmaLotada(indiceTurma) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "capacidade da sala insuficiente"})
		return
	}

	if alunoComConflitoDeHorario(indiceTurma, body.AlunoId) {
		c.JSON(http.StatusConflict, gin.H{"error": "conflito de horário com outra turma do aluno"})
		return
	}

	turmas[indiceTurma].Alunos = append(turmas[indiceTurma].Alunos, alunos[procurarAluno(body.AlunoId)])
	turmas[indiceTurma].QuantidadeAlunos = len(turmas[indiceTurma].Alunos)

	c.JSON(http.StatusOK, turmas[indiceTurma])
}

func listarAlunosDaTurma(c *gin.Context) {
	indice := buscarPorIdTurma(c)

	if indice == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrada"})
		return
	}

	c.JSON(http.StatusOK, turmas[indice].Alunos)
}

func alocarSala(c *gin.Context) {
	indiceTurma := buscarPorIdTurma(c)

	if indiceTurma == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma não encontrada"})
		return
	}

	var body struct {
		SalaId        int    `json:"sala_id"`
		DiaSemana     string `json:"dia_semana"`
		HorarioInicio string `json:"horario_inicio"`
		HorarioFim    string `json:"horario_fim"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if salaNaoExiste(body.SalaId) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrada"})
		return
	}

	if salaMenorQueTurma(indiceTurma, body.SalaId) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "capacidade da sala insuficiente"})
		return
	}

	if salaComConflitoDeHorario(indiceTurma, body.SalaId, body.DiaSemana, body.HorarioInicio, body.HorarioFim) {
		c.JSON(http.StatusConflict, gin.H{"error": "conflito de horário com outra turma na sala"})
		return
	}

	turmas[indiceTurma].SalaId = body.SalaId
	turmas[indiceTurma].DiaSemana = body.DiaSemana
	turmas[indiceTurma].HorarioInicio = body.HorarioInicio
	turmas[indiceTurma].HorarioFim = body.HorarioFim

	c.JSON(http.StatusOK, turmas[indiceTurma])
}

func buscarPorIdAluno(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1
	}

	for indice, aluno := range alunos {
		if aluno.Id == id {
			return indice
		}
	}

	return -1
}

func buscarPorIdSala(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1
	}

	for indice, sala := range salas {
		if sala.Id == id {
			return indice
		}
	}

	return -1
}

func buscarPorIdTurma(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1
	}

	for indice, turma := range turmas {
		if turma.Id == id {
			return indice
		}
	}

	return -1
}

func procurarAluno(id int) int {
	for indice, aluno := range alunos {
		if aluno.Id == id {
			return indice
		}
	}

	return -1
}

func procurarSala(id int) int {
	for indice, sala := range salas {
		if sala.Id == id {
			return indice
		}
	}

	return -1
}

func alunoEstaNaTurma(turma models.Turma, id int) bool {
	for _, aluno := range turma.Alunos {
		if aluno.Id == id {
			return true
		}
	}

	return false
}

func horariosSobrepoem(inicio1, fim1, inicio2, fim2 string) bool {
	hInicio1, _ := time.Parse("15:04", inicio1)
	hFim1, _ := time.Parse("15:04", fim1)
	hInicio2, _ := time.Parse("15:04", inicio2)
	hFim2, _ := time.Parse("15:04", fim2)

	return hInicio1.Before(hFim2) && hFim1.After(hInicio2)
}

func alunoNaoExiste(id int) bool {
	return procurarAluno(id) == -1
}

func salaNaoExiste(id int) bool {
	return procurarSala(id) == -1
}

func turmaLotada(indiceTurma int) bool {
	if turmas[indiceTurma].SalaId == 0 {
		return false
	}

	indiceSala := procurarSala(turmas[indiceTurma].SalaId)

	return indiceSala != -1 && len(turmas[indiceTurma].Alunos) >= salas[indiceSala].Capacidade
}

func alunoComConflitoDeHorario(indiceTurma, idAluno int) bool {
	if turmas[indiceTurma].SalaId == 0 {
		return false
	}

	for indice, turma := range turmas {
		if indice == indiceTurma || turma.SalaId == 0 {
			continue
		}

		if alunoEstaNaTurma(turma, idAluno) &&
			turma.DiaSemana == turmas[indiceTurma].DiaSemana &&
			horariosSobrepoem(turmas[indiceTurma].HorarioInicio, turmas[indiceTurma].HorarioFim, turma.HorarioInicio, turma.HorarioFim) {
			return true
		}
	}

	return false
}

func salaMenorQueTurma(indiceTurma, idSala int) bool {
	indiceSala := procurarSala(idSala)

	return salas[indiceSala].Capacidade < len(turmas[indiceTurma].Alunos)
}

func salaComConflitoDeHorario(indiceTurma, idSala int, dia, inicio, fim string) bool {
	for indice, turma := range turmas {
		if indice == indiceTurma || turma.SalaId != idSala {
			continue
		}

		if turma.DiaSemana == dia && horariosSobrepoem(inicio, fim, turma.HorarioInicio, turma.HorarioFim) {
			return true
		}
	}

	return false
}
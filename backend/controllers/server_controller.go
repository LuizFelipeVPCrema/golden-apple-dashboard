package controllers

import (
	"net/http"

	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/models"
	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/services"
	"github.com/gin-gonic/gin"
)

func CreateServer(c *gin.Context) {
	var req models.ServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	if err := services.CreateServer(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar servidor",
			"debug": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Servidor criado com sucesso!"})
}

func ListServers(c *gin.Context) {
	statusFilter := c.Query("status")
	projectFilter := c.Query("project")

	servers, err := services.ListServersWithFilters(statusFilter, projectFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao listar servidores",
			"debug": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, servers)
}

func DeleteServer(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome do servidor não fornecido."})
		return
	}

	err := services.DeleteServer(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao deletar servidor",
			"debug": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Servidor deletado com sucesso!"})

}

func StartServer(c *gin.Context) {
	name := c.Param("name")
	if err := services.StartServer(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao iniciar servidor", "debug": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Servidor iniciando com sucesso"})
}

func StopServer(c *gin.Context) {
	name := c.Param("name")
	if err := services.StopServer(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao fechar servidor", "debug": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Servidor parado com sucesso"})
}

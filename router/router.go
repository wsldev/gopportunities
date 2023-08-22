package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Iniciando o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a nossa api
	router.Run(":8080")
}

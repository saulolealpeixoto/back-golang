package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// isntancias de rotas do gin
	r := gin.Default()
	//rota para o endpoint /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "aplication is running",
		})
	})

	// porta que o servidor irá escutar
	r.Run(":8080")

}

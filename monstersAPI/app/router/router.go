package router

import (
	"battle-of-monsters/app/config"
	"battle-of-monsters/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	server := gin.Default()

	server.MaxMultipartMemory = config.ENV.MaxMemory

	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusText(http.StatusOK),
		})
	})

	// Battle routes
	battle := server.Group("/battle")
	{
		battle.GET("", controller.ListBattles)
	}

	// Monters routes
	monsters := server.Group("/monsters")
	{
		monsters.POST("", controller.CreateMonster)
		monsters.POST("/import", controller.ImportCSV)
		monsters.GET("/:monsterID", controller.FetchMonster)
		monsters.PUT("/:monsterID", controller.UpdateMonster)
		monsters.DELETE("/:monsterID", controller.DeleteMonster)
	}

	return server
}

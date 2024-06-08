package controller

import (
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBattles(context *gin.Context) {
	var battle []models.Battle

	var result *gorm.DB

	if result = db.CONN.Find(&battle); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Found %v battles", result.RowsAffected)
	context.JSON(http.StatusOK, &battle)
}

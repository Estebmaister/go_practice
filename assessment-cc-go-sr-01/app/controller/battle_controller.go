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

func CreateBattle(context *gin.Context) {
	var battleRequest struct {
		MonsterAID uint `json:"monster_a_id" binding:"required"`
		MonsterBID uint `json:"monster_b_id" binding:"required"`
	}

	if err := context.BindJSON(&battleRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monsterA := models.Monster{}
	err := fillMonsterByID(context, &monsterA, battleRequest.MonsterAID)
	if err != nil {
		return
	}

	monsterB := models.Monster{}
	err = fillMonsterByID(context, &monsterB, battleRequest.MonsterBID)
	if err != nil {
		return
	}

	battle := models.Battle{}
	battle.MonsterAID = battleRequest.MonsterAID
	battle.MonsterA = monsterA
	battle.MonsterBID = battleRequest.MonsterBID
	battle.MonsterB = monsterB

	battle.WinnerID = battleAction(monsterA, monsterB)
	if battle.WinnerID == monsterA.ID {
		battle.Winner = monsterA
	} else {
		battle.Winner = monsterB
	}

	if result := db.CONN.Create(&battle); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	log.Printf("battle %v has been created", battle.ID)

	context.JSON(http.StatusCreated, &battle)
}

func DeleteBattle(context *gin.Context) {
	battleID := context.Param("battleID")

	var battle models.Battle

	if result := db.CONN.First(&battle, battleID); result.Error != nil && result.Error.Error() != recordNotFound {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else if result.Error != nil && result.Error.Error() == recordNotFound {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})

		return
	}

	if result := db.CONN.Delete(&models.Battle{}, battleID); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	context.Status(http.StatusNoContent)
}

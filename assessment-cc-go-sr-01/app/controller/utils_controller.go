package controller

import (
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fillMonsterByID(context *gin.Context, monster *models.Monster, monsterID uint) error {
	if result := db.CONN.First(&monster, monsterID); result.Error != nil && result.Error.Error() != recordNotFound {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})

		return result.Error
	} else if result.Error != nil && result.Error.Error() == recordNotFound {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})

		return result.Error
	}

	return nil
}

func battleAction(monsterA, monsterB models.Monster) uint {
	var attacker1, attacker2 models.Monster

	// Deciding the order to attack
	if monsterA.Speed > monsterB.Speed {
		attacker1 = monsterA
		attacker2 = monsterB
	} else if monsterA.Speed == monsterB.Speed {
		if monsterA.Attack >= monsterB.Attack {
			attacker1 = monsterA
			attacker2 = monsterB
		} else {
			attacker1 = monsterB
			attacker2 = monsterA
		}
	} else {
		attacker1 = monsterB
		attacker2 = monsterA
	}

	for {
		// Round 1
		damage1 := calculateDamage(attacker1.Attack, attacker2.Defense)
		if damage1 >= int(attacker2.Hp) {
			return attacker1.ID
		}
		attacker2.Hp -= uint(damage1)

		// Round 2
		damage2 := calculateDamage(attacker2.Attack, attacker1.Defense)
		if damage2 >= int(attacker1.Hp) {
			return attacker2.ID
		}
		attacker1.Hp -= uint(damage2)
	}
}

func calculateDamage(attack, defense uint) int {
	var damage int = 1
	if attack > defense {
		damage = int(attack) - int(defense)
	}
	return damage
}

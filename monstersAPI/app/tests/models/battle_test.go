package models_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	utilstests "battle-of-monsters/app/tests/utils"
)

var _ = Describe("Battle", func() {
	utilstests.LoadEnv()
	db.Connect()

	BeforeEach(func() {
		if err := db.CONN.Exec("DELETE FROM battles; DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete battle and monsters. %w", err))
		}
	})

	AfterEach(func() {
		if err := db.CONN.Exec("DELETE FROM battles; DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete battle and monsters. %w", err))
		}
	})

	Describe("Marshal", func() {

		var b []byte
		var expected []byte

		JustBeforeEach(func() {
			blueSnake := models.Monster{
				Name:     "Blue Snake",
				Attack:   10,
				Defense:  15,
				Hp:       8,
				Speed:    18,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
			}

			redUnicorn := models.Monster{
				Name:     "Red Unicorn",
				Attack:   12,
				Defense:  14,
				Hp:       10,
				Speed:    9,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-unicorn.png",
			}

			db.CONN.Create(&blueSnake)
			db.CONN.Create(&redUnicorn)

			battle := models.Battle{
				MonsterA: models.Monster{
					ID:      blueSnake.ID,
					Name:    blueSnake.Name,
					Defense: blueSnake.Defense,
					Attack:  blueSnake.Attack,
					Hp:      blueSnake.Hp,
					Speed:   blueSnake.Speed,
				},
				MonsterB: models.Monster{
					ID:      redUnicorn.ID,
					Name:    redUnicorn.Name,
					Defense: redUnicorn.Defense,
					Attack:  redUnicorn.Attack,
					Hp:      redUnicorn.Hp,
					Speed:   redUnicorn.Speed,
				},
				Winner: models.Monster{
					ID:      blueSnake.ID,
					Name:    blueSnake.Name,
					Defense: blueSnake.Defense,
					Attack:  blueSnake.Attack,
					Hp:      blueSnake.Hp,
					Speed:   blueSnake.Speed,
				},
			}

			if r := db.CONN.Create(&battle); r.Error != nil {
				log.Fatalln(r.Error)
			}

			db.CONN.Preload("Winner").Preload("MonsterA").Preload("MonsterB").First(&battle, battle.ID)

			b, _ = json.Marshal(battle)

			expected = []byte(`{
				"id": 1,
				"monsterA": {
					"id": 1,
          "name": "Blue Snake",
          "attack": 10,
          "defense": 15,
          "hp": 8,
          "speed": 18,
          "imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
          "battles": null
				},
				"monsterB": {
					"id": 2,
          "name": "Red Unicorn",
          "attack": 12,
          "defense": 14,
          "hp": 10,
          "speed": 9,
          "imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-unicorn.png",
          "battles": null
				},
				"winner": {
					"id": 1,
          "name": "Blue Snake",
          "attack": 10,
          "defense": 15,
          "hp": 8,
          "speed": 18,
          "imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
          "battles": null
				}
			}`)

		})

		Context("should parse battle correctly", func() {

			It("battle should match with the expected json", func() {
				Expect(b).Should(MatchJSON(expected))
			})

		})

	})

	Describe("Validate", func() {

		When("creating a battle", func() {

			var errModel error

			JustBeforeEach(func() {
				blueSnake := models.Monster{
					Name:     "Blue Snake",
					Attack:   10,
					Defense:  15,
					Hp:       8,
					Speed:    18,
					ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
				}

				db.CONN.Create(&blueSnake)

				battle := models.Battle{
					MonsterA: models.Monster{
						ID:      blueSnake.ID,
						Name:    blueSnake.Name,
						Defense: blueSnake.Defense,
						Attack:  blueSnake.Attack,
						Hp:      blueSnake.Hp,
						Speed:   blueSnake.Speed,
					},
					Winner: models.Monster{
						ID:      blueSnake.ID,
						Name:    blueSnake.Name,
						Defense: blueSnake.Defense,
						Attack:  blueSnake.Attack,
						Hp:      blueSnake.Hp,
						Speed:   blueSnake.Speed,
					},
				}

				if r := db.CONN.Create(&battle); r.Error != nil {
					errModel = r.Error
				}

			})

			Context("should validate battle without monsterB", func() {

				It("should return a error", func() {
					var err models.ValidationErrors
					Expect(errors.As(errModel, &err)).To(BeTrue())
				})

			})

		})

	})

})

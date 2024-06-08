package models_test

import (
	"encoding/json"
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	utilstests "battle-of-monsters/app/tests/utils"
)

var _ = Describe("Monster", func() {
	utilstests.LoadEnv()
	db.Connect()

	BeforeEach(func() {
		if err := db.CONN.Exec("DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}
	})

	AfterEach(func() {
		if err := db.CONN.Exec("DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}
	})

	Describe("Marshal", func() {

		var darkSnake *models.Monster
		var m []byte
		var expected []byte

		JustBeforeEach(func() {
			darkSnake = &models.Monster{
				Name:     "Dark Snake",
				Attack:   10,
				Defense:  15,
				Hp:       8,
				Speed:    18,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dark-snake.png",
			}

			db.CONN.Create(darkSnake)

			m, _ = json.Marshal(darkSnake)

			expected = []byte(`{
				"id": 1,
				"name": "Dark Snake",
				"attack": 10,
				"defense": 15,
				"hp": 8,
				"speed": 18,
				"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dark-snake.png",
				"battles": null
			}`)

		})

		Context("should parse moster correctly", func() {

			It("monster should match with the expected json", func() {
				Expect(m).Should(MatchJSON(expected))
			})

		})

	})

	Describe("Validate", func() {

		When("creating a monster", func() {

			var monster *models.Monster
			var errModel error

			JustBeforeEach(func() {
				monster = &models.Monster{
					Attack:   10,
					Defense:  15,
					Hp:       8,
					Speed:    18,
					ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/no-named.png",
				}

				if r := db.CONN.Create(monster); r.Error != nil {
					errModel = r.Error
				}

			})

			Context("should validate monster without name", func() {

				It("should return a error", func() {
					var err models.ValidationErrors
					Expect(errors.As(errModel, &err)).To(BeTrue())
				})

			})

		})

		When("updating a monster", func() {

			var monster *models.Monster
			var errModel error

			JustBeforeEach(func() {
				monster = &models.Monster{
					Name:     "Atomic Robot",
					Attack:   11,
					Defense:  12,
					Hp:       7,
					Speed:    11,
					ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/atomic-robot.png",
				}

				db.CONN.Create(monster)

				monster.Name = ""

				if r := db.CONN.Save(&monster); r.Error != nil {
					errModel = r.Error
				}

			})

			Context("should validate monster without name", func() {

				It("should return a error", func() {
					var err models.ValidationErrors
					Expect(errors.As(errModel, &err)).To(BeTrue())
				})

			})

		})

	})

})

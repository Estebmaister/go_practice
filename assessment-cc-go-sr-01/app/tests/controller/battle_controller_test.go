package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	utilstests "battle-of-monsters/app/tests/utils"
)

var _ = Describe("BattleController", func() {
	utilstests.LoadEnv()
	db.Connect()

	var blueSnake *models.Monster
	var redUnicorn *models.Monster
	var insectSnake *models.Monster
	var battle *models.Battle

	BeforeEach(func() {
		if err := db.CONN.Exec("DELETE FROM battles; DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete battle and monsters. %w", err))
		}

		blueSnake = &models.Monster{
			Name:     "Blue Snake",
			Attack:   10,
			Defense:  15,
			Hp:       8,
			Speed:    18,
			ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
		}

		redUnicorn = &models.Monster{
			Name:     "Red Unicorn",
			Attack:   12,
			Defense:  14,
			Hp:       10,
			Speed:    9,
			ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-unicorn.png",
		}

		insectSnake = &models.Monster{
			Name:     "Insect Snake",
			Attack:   87,
			Defense:  2,
			Hp:       6,
			Speed:    17,
			ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-unicorn.png",
		}

		db.CONN.Create(blueSnake)
		db.CONN.Create(redUnicorn)
		db.CONN.Create(insectSnake)

		battle = &models.Battle{
			MonsterA: *blueSnake,
			MonsterB: *redUnicorn,
			Winner:   *blueSnake,
		}

		db.CONN.Create(battle)
	})

	AfterEach(func() {
		if err := db.CONN.Exec("DELETE FROM battles; DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete battle and monsters. %w", err))
		}
	})

	Describe("List", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodGet, "/battle", nil)
			response = utilstests.ExecuteRequest(req)
		})

		Context("should list all battles", func() {

			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				l, _ := utilstests.DeserializeList(response.Body.String())
				Expect(len(l)).Should(BeNumerically(">=", 0))
			})

		})

	})

	Describe("Battle", func() {
		var response *httptest.ResponseRecorder
		var payload []byte

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodPost, "/battle", bytes.NewBuffer(payload))
			response = utilstests.ExecuteRequest(req)
		})

		Context("should fail when trying a battle of monsters with an undefined monster", func() {

			BeforeEach(func() {
				payload = []byte(`{"monster_a_id": 1}`)
			})

			It("status code should be 400", func() {
				Expect(response.Code).To(Equal(400))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

		Context("should fail when trying a battle of monsters with an inexistent monster", func() {

			BeforeEach(func() {
				payload = []byte(`{"monster_a_id": 1, "monster_b_id": 999}`)
			})

			It("status code should be 404", func() {
				Expect(response.Code).To(Equal(404))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

		Context("should insert a battle of monsters successfully with monster 1 winning", func() {

			BeforeEach(func() {
				payload = []byte(`{"monster_a_id": 1, "monster_b_id": 3}`)
			})

			It("status code should be 201", func() {
				Expect(response.Code).To(Equal(201))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				battle := models.Battle{}
				if err := json.Unmarshal(response.Body.Bytes(), &battle); err != nil {
					Expect(err).Error().To(BeNil())
				}

				Expect(battle.Winner.ID).Should(BeEquivalentTo(1))
			})

		})

		Context("should insert a battle of monsters successfully with monster 2 winning", func() {

			BeforeEach(func() {
				payload = []byte(`{"monster_a_id": 2, "monster_b_id": 1}`)
			})

			It("status code should be 201", func() {
				Expect(response.Code).To(Equal(201))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				battle := models.Battle{}
				if err := json.Unmarshal(response.Body.Bytes(), &battle); err != nil {
					Expect(err).Error().To(BeNil())
				}
				Expect(battle.Winner.ID).Should(BeEquivalentTo(2))
			})

		})

	})

	Describe("Delete", func() {
		var response *httptest.ResponseRecorder
		var dragon *models.Monster
		var robot *models.Monster
		var dr *models.Battle

		BeforeEach(func() {
			dragon = &models.Monster{
				Name:     "Dragon",
				Attack:   10,
				Defense:  15,
				Hp:       18,
				Speed:    5,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dragon.png",
			}

			robot = &models.Monster{
				Name:     "Robot",
				Attack:   12,
				Defense:  13,
				Hp:       15,
				Speed:    7,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/robot.png",
			}

			db.CONN.Create(dragon)
			db.CONN.Create(robot)

			dr = &models.Battle{
				MonsterA: *dragon,
				MonsterB: *robot,
				Winner:   *dragon,
			}

			db.CONN.Create(dr)
		})

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodDelete, "/battle/"+fmt.Sprintf("%v", dr.ID), nil)
			response = utilstests.ExecuteRequest(req)
		})

		Context("should delete a battle correctly", func() {

			It("status code should be 204", func() {
				Expect(response.Code).To(Equal(204))
			})

		})

		Context("should get a 404 response", func() {

			JustBeforeEach(func() {
				req, _ := http.NewRequest(http.MethodDelete, "/battle/999", nil)
				response = utilstests.ExecuteRequest(req)
			})

			It("status code should be 404", func() {
				Expect(response.Code).To(Equal(404))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

	})

})

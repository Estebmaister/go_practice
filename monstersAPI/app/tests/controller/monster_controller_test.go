package controller_test

import (
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	utilstests "battle-of-monsters/app/tests/utils"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MonsterController", func() {
	utilstests.LoadEnv()
	db.Connect()

	var monster *models.Monster

	BeforeEach(func() {
		if err := db.CONN.Exec("DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}

		monster = &models.Monster{
			Name:     "Blue Snake",
			Attack:   10,
			Defense:  15,
			Hp:       8,
			Speed:    18,
			ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
		}

		db.CONN.Create(monster)
	})

	AfterEach(func() {
		if err := db.CONN.Exec("DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}
	})

	Describe("List", func() {
		var response *httptest.ResponseRecorder
		var expectedURL = "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png"

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodGet, "/monsters", nil)
			response = utilstests.ExecuteRequest(req)
		})

		Context("should list all monsters", func() {

			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				l, _ := utilstests.DeserializeList(response.Body.String())
				Expect(len(l)).To(Equal(1))
				for _, m := range l {
					Expect(m["id"]).To(BeEquivalentTo(monster.ID))
					Expect(m["name"]).To(Equal("Blue Snake"))
					Expect(m["attack"]).To(BeEquivalentTo(10))
					Expect(m["defense"]).To(BeEquivalentTo(15))
					Expect(m["hp"]).To(BeEquivalentTo(8))
					Expect(m["speed"]).To(BeEquivalentTo(18))
					Expect(m["imageUrl"]).To(Equal(expectedURL))
				}
			})

		})

	})

	Describe("Get", func() {
		var response *httptest.ResponseRecorder
		var expectedURL = "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png"

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodGet, "/monsters/"+fmt.Sprintf("%v", monster.ID), nil)
			response = utilstests.ExecuteRequest(req)
		})

		Context("should get a monster correctly", func() {

			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				m, _ := utilstests.Deserialize(response.Body.String())
				Expect(m["id"]).To(BeEquivalentTo(monster.ID))
				Expect(m["name"]).To(Equal("Blue Snake"))
				Expect(m["attack"]).To(BeEquivalentTo(10))
				Expect(m["defense"]).To(BeEquivalentTo(15))
				Expect(m["hp"]).To(BeEquivalentTo(8))
				Expect(m["speed"]).To(BeEquivalentTo(18))
				Expect(m["imageUrl"]).To(Equal(expectedURL))
			})

		})

		Context("should get a 404 response", func() {

			JustBeforeEach(func() {
				req, _ := http.NewRequest(http.MethodGet, "/monsters/0", nil)
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

	Describe("Create", func() {
		var response *httptest.ResponseRecorder
		var payload []byte
		var expectedURL = "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-dragon.png"

		BeforeEach(func() {
			payload = []byte(`{
				"name": "Red Dragon",
				"attack": 10, 
				"defense": 10, 
				"hp": 10, 
				"speed": 10, 
				"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-dragon.png"
			}`)
		})

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodPost, "/monsters", bytes.NewBuffer(payload))
			response = utilstests.ExecuteRequest(req)
		})

		Context("should create a monster correctly", func() {

			It("status code should be 201", func() {
				Expect(response.Code).To(Equal(201))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				m, _ := utilstests.Deserialize(response.Body.String())
				Expect(m["id"]).ToNot(BeNil())
				Expect(m["name"]).To(Equal("Red Dragon"))
				Expect(m["attack"]).To(BeEquivalentTo(10))
				Expect(m["defense"]).To(BeEquivalentTo(10))
				Expect(m["hp"]).To(BeEquivalentTo(10))
				Expect(m["speed"]).To(BeEquivalentTo(10))
				Expect(m["imageUrl"]).To(Equal(expectedURL))
			})

		})

		Context("should get an error when create a monster without name", func() {

			BeforeEach(func() {
				payload = []byte(`{
					"attack": 10, 
					"defense": 10, 
					"hp": 10, 
					"speed": 10, 
					"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-dragon.png"
				}`)
			})

			It("status code should be 400", func() {
				Expect(response.Code).To(Equal(400))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

	})

	Describe("Update", func() {
		var response *httptest.ResponseRecorder
		var payload []byte
		var payloadID uint
		var expectedURL = "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dark-snake.png"

		BeforeEach(func() {
			payloadID = monster.ID
			payload = []byte(`{
				"id": 1,
				"name": "Dark Snake",
				"attack": 10, 
				"defense": 10, 
				"hp": 10, 
				"speed": 10, 
				"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dark-snake.png"
			}`)
		})

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodPut, "/monsters/"+fmt.Sprintf("%v", payloadID), bytes.NewBuffer(payload))
			response = utilstests.ExecuteRequest(req)
		})

		Context("should update a monster correctly", func() {

			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				m, _ := utilstests.Deserialize(response.Body.String())
				Expect(m["id"]).To(BeEquivalentTo(1))
				Expect(m["name"]).To(Equal("Dark Snake"))
				Expect(m["attack"]).To(BeEquivalentTo(10))
				Expect(m["defense"]).To(BeEquivalentTo(10))
				Expect(m["hp"]).To(BeEquivalentTo(10))
				Expect(m["speed"]).To(BeEquivalentTo(10))
				Expect(m["imageUrl"]).To(Equal(expectedURL))
			})

		})

		Context("should get an error when update a monster without name", func() {

			BeforeEach(func() {
				payloadID = monster.ID
				payload = []byte(`{
					"attack": 10, 
					"defense": 10, 
					"hp": 10, 
					"speed": 10, 
					"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-dragon.png"
				}`)
			})

			It("status code should be 400", func() {
				Expect(response.Code).To(Equal(400))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

		Context("should get an error when update a monster that does not exists", func() {

			BeforeEach(func() {
				payloadID = 999
				payload = []byte(`{
					"id": 999,
					"name": "Dead Snake",
					"attack": 16, 
					"defense": 1, 
					"hp": 13, 
					"speed": 12, 
					"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dead-snake.png"
				}`)
			})

			It("status code should be 404", func() {
				Expect(response.Code).To(Equal(404))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

		})

	})

	Describe("Delete", func() {
		var response *httptest.ResponseRecorder
		var dragon *models.Monster

		BeforeEach(func() {
			dragon = &models.Monster{
				Name:     "Dragon",
				Attack:   10,
				Defense:  15,
				Hp:       18,
				Speed:    5,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dragon.png",
			}

			db.CONN.Create(dragon)
		})

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodDelete, "/monsters/"+fmt.Sprintf("%v", dragon.ID), nil)
			response = utilstests.ExecuteRequest(req)
		})

		Context("should delete a monster correctly", func() {

			It("status code should be 204", func() {
				Expect(response.Code).To(Equal(204))
			})

		})

		Context("should get a 404 response", func() {

			JustBeforeEach(func() {
				req, _ := http.NewRequest(http.MethodDelete, "/monsters/999", nil)
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

	// TODO Implement the tests below
	Describe("Import CSV", func() {
		var _ *httptest.ResponseRecorder

		JustBeforeEach(func() {
			req, _ := http.NewRequest(http.MethodPost, "/monsters/import", nil)
			_ = utilstests.ExecuteRequest(req)
		})

		Context("should fail when importing csv file with an empty monster", func() {

			PIt("status code should be 400")

		})

		Context("should fail when importing csv file with wrong or inexistent columns.", func() {

			PIt("status code should be 400")

		})

		Context("should import all the CSV objects into the database successfully", func() {

			PIt("status code should be 200")

		})

	})

})

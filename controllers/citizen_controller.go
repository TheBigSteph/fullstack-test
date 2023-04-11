package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TheBigSteph/fullstack-test/configs"
	"github.com/TheBigSteph/fullstack-test/models"
	"github.com/TheBigSteph/fullstack-test/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var citizenCollection *mongo.Collection = configs.GetCollection(configs.DB, "citizens")
var validate = validator.New()

func CreateCitizen(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var citizen models.Citizen
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&citizen); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CitizenResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&citizen); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CitizenResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	
	newCitizen := models.Citizen{
		Id:         primitive.NewObjectID(),
		Area:       citizen.Area,
		Address:    citizen.Address,
		Profession: citizen.Profession,
		Gender:     citizen.Gender,
		User: models.User{
			FirstName: citizen.User.FirstName,
			LastName:  citizen.User.LastName,
			Email:     citizen.User.Email,
		},
		Requests: append(citizen.Requests, models.Requests{}),
	}

	result, err := citizenCollection.InsertOne(ctx, newCitizen)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CitizenResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.CitizenResponse{Status: http.StatusCreated, Message: "created", Data: &fiber.Map{"data": result}})
}

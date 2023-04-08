package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	Id primitive.ObjectID `json:"id.omitempty"`
	Subject string `json:"subject.omitempty" validate:"required"`
	Description string `json:"description"`
	Status string `json:"status"`
	Citizen Citizen `json:"citizen"`
}
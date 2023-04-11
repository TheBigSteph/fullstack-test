package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	) 
	

type User struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type Requests struct {
	Id primitive.ObjectID `json:"id.omitempty"`
	Subject string `json:"subject" validate:"required"`
	Description string `json:"description"`
	Status string `json:"status"`
}

type Citizen struct {
	Id primitive.ObjectID `json:"id.omitempty"`
	Area string `json:"area" validate:"required"`
	Address string `json:"address" validate:"required"`
	Profession string `json:"profession" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	User User `json:"user"`
	Requests []Requests `json:"requests"`
}

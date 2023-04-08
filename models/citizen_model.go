package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Citizen struct {
	Id primitive.ObjectID `json:"id.omitempty"`
	Area string `json:"area"`
	Address string `json:"address"`
	Profession string `json:"profession"`
	Gender string `json:"gender"`
	User []User `json:"user"`
}

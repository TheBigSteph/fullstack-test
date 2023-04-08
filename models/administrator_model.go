package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Administrator struct {
	Id primitive.ObjectID `json:"id.omitempty"`
	Department string `json:"department"`
}
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	Id   primitive.ObjectID `json:"_id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required"`
	Url  string             `json:"url,omitempty" validate:"required"`
}

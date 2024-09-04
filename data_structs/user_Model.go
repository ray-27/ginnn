package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `bson: "_id"`
	First_Name *string            `json: "first_name" validate:"required, min=2, max=100"`
}

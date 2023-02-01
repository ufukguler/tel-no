package model_entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CrawledProjection struct {
	ID          primitive.ObjectID `bson:"_id"`
	PhoneNumber string             `bson:"phoneNumber"`
	Comments    struct {
		ID          primitive.ObjectID `bson:"_id"`
		Comment     string             `bson:"comment"`
		CreatedDate time.Time          `bson:"createdDate"`
		Updated     bool               `bson:"updated"`
	} `bson:"comments"`
}

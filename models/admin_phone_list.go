package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CrawledUrlAdmin struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
	Active      bool               `bson:"active" json:"active"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate" json:"updatedDate"`
}

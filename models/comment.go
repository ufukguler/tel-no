package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SingleComment struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Comment     string             `bson:"comment" json:"comment"`
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
	CommentType string             `bson:"commentType" json:"commentType"`
	Updated     bool               `bson:"updated" json:"updated"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate" json:"updatedDate"`
}

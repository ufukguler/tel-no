package model_entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CrawledUrl struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	PhoneNumber   string             `bson:"phoneNumber" json:"phoneNumber"`
	Comments      []Comment          `bson:"comments" json:"comments"`
	Active        bool               `bson:"active" json:"active"`
	ViewCount     int                `bson:"reviewCount" json:"reviewCount"`
	LastRequestIP string             `bson:"lastRequestIP" json:"lastRequestIP"`
	CreatedDate   time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate   time.Time          `bson:"updatedDate" json:"updatedDate"`
}

type Comment struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Comment     string             `bson:"comment" json:"comment"`
	CommentType string             `bson:"commentType" json:"commentType"`
	Updated     bool               `bson:"updated" json:"updated"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate" json:"updatedDate"`
}

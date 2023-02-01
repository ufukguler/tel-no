package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"telno/database"
	"telno/model_entity"
	"time"
)

const (
	crawledUrlCommentId = "comments._id"
)

func FindCommentById(id string) (string, model_entity.Comment, error) {
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		return "", model_entity.Comment{}, err2
	}

	var crawledUrl model_entity.CrawledUrl
	filter := bson.M{crawledUrlCommentId: bson.M{"$in": []primitive.ObjectID{hex}}}
	if err := database.Database.FindOne(ColCrawledUrl, filter, &crawledUrl); err != nil {
		return "", model_entity.Comment{}, err
	}
	var comment model_entity.Comment
	for i := range crawledUrl.Comments {
		if crawledUrl.Comments[i].Id.Hex() == id {
			comment = crawledUrl.Comments[i]
		}
	}
	return crawledUrl.PhoneNumber, comment, nil
}

func DeleteCommentById(id string) error {
	// FIND COMMENT
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		return err2
	}

	var crawledUrl model_entity.CrawledUrl
	filter := bson.M{crawledUrlCommentId: bson.M{"$in": []primitive.ObjectID{hex}}}
	if err := database.Database.FindOne(ColCrawledUrl, filter, &crawledUrl); err != nil {
		return err
	}

	// SET COMMENTS OF CRAWLED URL
	comments := make([]model_entity.Comment, 0)
	for i := range crawledUrl.Comments {
		if crawledUrl.Comments[i].Id.Hex() != id {
			comments = append(comments, crawledUrl.Comments[i])
		}
	}

	// UPDATE COMMENTS OF CRAWLED URL
	crawledUrl.Comments = comments
	filter = bson.M{"_id": crawledUrl.Id}
	updateStatement := bson.M{"$set": crawledUrl}
	return database.Database.UpdateOne(ColCrawledUrl, filter, updateStatement)
}

func UpdateCommentById(id, comment, commentType string) error {
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		return err2
	}

	var crawledUrl model_entity.CrawledUrl
	filter := bson.M{crawledUrlCommentId: bson.M{"$in": []primitive.ObjectID{hex}}}
	if err := database.Database.FindOne(ColCrawledUrl, filter, &crawledUrl); err != nil {
		return err
	}
	for i := range crawledUrl.Comments {
		if crawledUrl.Comments[i].Id.Hex() == id {
			crawledUrl.Comments[i].Updated = true
			crawledUrl.Comments[i].Comment = comment
			crawledUrl.Comments[i].CommentType = commentType
			crawledUrl.Comments[i].UpdatedDate = time.Now()
			crawledUrl.UpdatedDate = time.Now()
		}
	}

	filter = bson.M{_id: crawledUrl.Id}
	updateStatement := bson.M{"$set": crawledUrl}

	return database.Database.UpdateOne(ColCrawledUrl, filter, updateStatement)
}

func QuickUpdateCommentById(id string) error {
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		return err2
	}

	var crawledUrl model_entity.CrawledUrl
	filter := bson.M{crawledUrlCommentId: bson.M{"$in": []primitive.ObjectID{hex}}}
	if err := database.Database.FindOne(ColCrawledUrl, filter, &crawledUrl); err != nil {
		return err
	}
	for i := range crawledUrl.Comments {
		if crawledUrl.Comments[i].Id.Hex() == id {
			crawledUrl.Comments[i].Updated = true
		}
	}

	filter = bson.M{_id: crawledUrl.Id}
	updateStatement := bson.M{"$set": crawledUrl}
	return database.Database.UpdateOne(ColCrawledUrl, filter, updateStatement)
}

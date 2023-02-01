package service

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"telno/database"
)

// const crawledUrlCommentUpdate string = "comments.updated"
const phoneNumber string = "phoneNumber"

func FindPhoneNumberPageable(limit, page int64, number string) (*mongopagination.PaginatedData, error) {
	filter := bson.M{"$match": bson.M{"$and": []bson.M{
		{phoneNumber: bson.M{"$regex": ".*" + number + ".*"}},
	}}}
	filter2 := bson.M{"$sort": bson.M{"updatedDate": -1}}
	return database.Database.FindByPageableMatch(ColCrawledUrl, limit, page, filter, filter2)
}

func FindPhoneNumberByCommentUncheckedPageable(limit, page int64, number string) (*mongopagination.PaginatedData, error) {
	filter := bson.M{"$match": bson.M{"$and": []bson.M{
		{phoneNumber: bson.M{"$regex": ".*" + number + ".*"}},
		//{crawledUrlCommentUpdate: bson.M{"$ne": true}},
	}}}
	filter2 := bson.M{"$sort": bson.M{"updatedDate": -1}}
	return database.Database.FindByPageableMatch(ColCrawledUrl, limit, page, filter, filter2)
}

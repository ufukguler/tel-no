package service

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"telno/database"
	"telno/model_entity"
	"time"
)

const (
	ColCrawledUrl         = "crawled-url"
	crawledUrlPhoneNumber = "phoneNumber"
	_id                   = "_id"
	lastRequestIP         = "lastRequestIP"
	match                 = "$match"
	limit                 = "$limit"
	sample                = "$sample"
	ne                    = "$ne"
)

func FindByPhoneNumber(phoneNumber string) (model_entity.CrawledUrl, error) {
	phoneNumber = NormalizePhoneNumber(phoneNumber)

	var crawledUrl model_entity.CrawledUrl
	filter := bson.M{crawledUrlPhoneNumber: phoneNumber}
	if err := database.Database.FindOne(ColCrawledUrl, filter, &crawledUrl); err != nil {
		return model_entity.CrawledUrl{}, err
	}
	return crawledUrl, nil
}

func FindAllForSitemap() ([]model_entity.CrawledUrl, error) {
	var crawledUrl []model_entity.CrawledUrl

	if err := database.Database.Find(ColCrawledUrl, bson.M{}, &crawledUrl); err != nil {
		return []model_entity.CrawledUrl{}, err
	}
	return crawledUrl, nil
}

func NormalizePhoneNumber(phoneNumber string) string {
	phoneNumber = strings.TrimSpace(phoneNumber)
	if len(phoneNumber) == 0 {
		return ""
	}

	runes := []rune(phoneNumber)

	if phoneNumber[:3] == "444" && len(phoneNumber) == 7 {
		return phoneNumber
	}
	if string(runes[0]) == "+" {
		runes = runes[1:]
	}
	if string(runes[0]) == "9" {
		runes = runes[1:]
	}
	if string(runes[0]) == "0" {
		runes = runes[1:]
	}
	return string(runes)
}

func UpdatePhoneNumber(crawledUrl model_entity.CrawledUrl) error {
	filter := bson.M{_id: crawledUrl.Id}
	crawledUrl.UpdatedDate = time.Now()
	updateStatement := bson.M{"$set": crawledUrl}
	return database.Database.UpdateOne(ColCrawledUrl, filter, updateStatement)
}

func CreatePhoneNumber(crawledUrl model_entity.CrawledUrl) error {
	return database.Database.InsertOne(ColCrawledUrl, crawledUrl)
}

func FindSimilarNumbers(phoneNumber string) ([]model_entity.CrawledUrl, error) {
	var similarPhoneNumbers []model_entity.CrawledUrl
	filter := []bson.M{
		{match: bson.M{crawledUrlPhoneNumber: bson.M{"$regex": "^" + phoneNumber[:3] + "*"}}},
		{limit: 6},
	}
	err := database.Database.AggregateQuery(ColCrawledUrl, filter, &similarPhoneNumbers)
	if len(similarPhoneNumbers) < 5 {
		numbers := findRandomNumbers()
		max := 5 - len(similarPhoneNumbers)
		for _, v := range numbers[:max] {
			similarPhoneNumbers = append(similarPhoneNumbers, v)
		}
	}
	return similarPhoneNumbers, err
}

func findRandomNumbers() []model_entity.CrawledUrl {
	var similarPhoneNumbers []model_entity.CrawledUrl
	filter := []bson.M{
		{sample: bson.M{"size": 5}},
	}
	err := database.Database.AggregateQuery(ColCrawledUrl, filter, &similarPhoneNumbers)
	if err != nil {
		similarPhoneNumbers = make([]model_entity.CrawledUrl, 0)
	}
	return similarPhoneNumbers
}

func UpdateRequestCount(id primitive.ObjectID, remoteAddr string) {
	var crawledUrl model_entity.CrawledUrl
	filter := bson.M{
		_id:           id,
		lastRequestIP: bson.M{ne: remoteAddr},
	}
	err := database.Database.FindOne(ColCrawledUrl, filter, &crawledUrl)
	if err != nil {
		log.Errorf("(UpdateRequestCount) could not find phoneNumber by ID %s: and lastRemoteAdd: %s -- is user spamming?", id.Hex(), remoteAddr)
		return
	}
	crawledUrl.ViewCount = crawledUrl.ViewCount + 1
	crawledUrl.LastRequestIP = remoteAddr
	updateStatement := bson.M{"$set": crawledUrl}
	err = database.Database.UpdateOne(ColCrawledUrl, filter, updateStatement)
	if err != nil {
		log.Errorf("ERR (UpdateRequestCount) could not update phoneNumber by id %s", id.Hex())
		return
	}
}

func commentsUpdatedToday() []model_entity.CrawledUrl {
	aDayAgo := time.Now().Add(-24 * time.Hour)

	crawledUrl := make([]model_entity.CrawledUrl, 0)
	filter := bson.M{
		"updatedDate": bson.M{
			"$gte": aDayAgo,
		},
	}
	err := database.Database.Find(ColCrawledUrl, filter, &crawledUrl)
	if err != nil {
		log.Errorf("Error finding last updated numbers: %s", err.Error())
	}
	return crawledUrl
}

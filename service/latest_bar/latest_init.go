package latest_bar

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
	"telno/config"
	"telno/database"
	"telno/model_entity"
	"telno/models"
	"telno/service"
	"time"
)

func InitLatestSearches() {
	var sample []model_entity.CrawledUrl
	filter := []bson.M{
		{"$sample": bson.M{"size": 10}},
	}

	if err := database.Database.AggregateQuery(service.ColCrawledUrl, filter, &sample); err == nil {
		for i, k := range sample {
			if i == 10 {
				break
			}
			time.Sleep(time.Millisecond)
			config.ApiConfig.LatestSearchesMap[k.PhoneNumber] = models.LatestSearches{
				PhoneNumber: k.PhoneNumber,
				AddedAt:     time.Now().UnixMicro(),
			}
		}
	}
}

func InitLatestComments() {
	var sample []model_entity.CrawledProjection
	filter := []bson.M{
		{"$unwind": bson.M{"path": "$comments"}},
		{"$project": bson.M{
			"_id":                  0,
			"comments._id":         1,
			"phoneNumber":          1,
			"comments.updated":     1,
			"comments.comment":     1,
			"comments.createdDate": 1,
		}},
		{"$match": bson.M{"comments.updated": true}},
		{"$sort": bson.M{"comments.createdDate": -1}},
		{"$limit": 10},
	}

	if err := database.Database.AggregateQuery(service.ColCrawledUrl, filter, &sample); err == nil {
		sort.Slice(sample, func(i, j int) bool {
			return sample[i].Comments.CreatedDate.After(sample[j].Comments.CreatedDate)
		})
		config.ApiConfig.LatestCommentsMap = make(map[string]models.LatestComments, 0)
		for i, _ := range sample {
			if sample[i].Comments.Updated {
				config.ApiConfig.LatestCommentsMap[sample[i].Comments.Comment] = models.LatestComments{
					PhoneNumber: sample[i].PhoneNumber,
					Comment:     sample[i].Comments.Comment,
					AddedAt:     sample[i].Comments.CreatedDate.UnixMilli(),
				}
			}
			if len(config.ApiConfig.LatestCommentsMap) == 10 {
				break
			}
		}
	} else {
		log.Errorf("aggreagate err: %v", err.Error())
	}
}

package latest_bar

import (
	"sort"
	"telno/config"
	"telno/models"
	"time"
)

func GetLatestComments() []models.LatestComments {
	comments := make([]models.LatestComments, 0)
	for s := range config.ApiConfig.LatestCommentsMap {
		comments = append(comments, config.ApiConfig.LatestCommentsMap[s])
		if len(comments) == 10 {
			break
		}
	}
	sortLatestComments(comments)
	return comments
}

func UpdateLatestComments(phoneNumber, comment string) {
	InitLatestComments()

	if false { // todo
		comments := make([]models.LatestComments, 0)
		for s := range config.ApiConfig.LatestCommentsMap {
			comments = append(comments, config.ApiConfig.LatestCommentsMap[s])
		}
		_, ok := config.ApiConfig.LatestCommentsMap[comment]
		if ok {
			latestComments := config.ApiConfig.LatestCommentsMap[comment]
			latestComments.AddedAt = getNowMicro()
			config.ApiConfig.LatestCommentsMap[comment] = latestComments
		} else {
			latestComments := models.LatestComments{
				PhoneNumber: phoneNumber,
				Comment:     comment,
				AddedAt:     getNowMicro(),
			}
			comments = append(comments, latestComments)
			config.ApiConfig.LatestCommentsMap[comment] = latestComments
		}
		sortLatestComments(comments)
		commentsUpdate := comments[:getMax(len(comments))]
		config.ApiConfig.LatestCommentsMap = make(map[string]models.LatestComments, 0)
		for _, latestComments := range commentsUpdate {
			config.ApiConfig.LatestCommentsMap[latestComments.PhoneNumber] = models.LatestComments{
				PhoneNumber: latestComments.PhoneNumber,
				AddedAt:     latestComments.AddedAt,
			}
		}
	}
}

func getNowMicro() int64 {
	return time.Now().UnixMicro()
}

func sortLatestComments(searches []models.LatestComments) {
	sort.Slice(searches, func(i, j int) bool {
		return searches[i].AddedAt > searches[j].AddedAt
	})
}

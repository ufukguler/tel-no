package latest_bar

import (
	"sort"
	"telno/config"
	"telno/models"
)

func GetLatestSearches() []models.LatestSearches {
	searches := make([]models.LatestSearches, 0)
	for s := range config.ApiConfig.LatestSearchesMap {
		searches = append(searches, config.ApiConfig.LatestSearchesMap[s])
		if len(searches) == 10 {
			break
		}
	}
	sortLatestSearches(searches)
	return searches
}

func UpdateLatestSearches(phoneNumber string) {
	searches := make([]models.LatestSearches, 0)
	for s := range config.ApiConfig.LatestSearchesMap {
		searches = append(searches, config.ApiConfig.LatestSearchesMap[s])
	}
	_, ok := config.ApiConfig.LatestSearchesMap[phoneNumber]
	if ok {
		latestSearches := config.ApiConfig.LatestSearchesMap[phoneNumber]
		latestSearches.AddedAt = getNowMicro()
		config.ApiConfig.LatestSearchesMap[phoneNumber] = latestSearches
	} else {
		latestSearches := models.LatestSearches{
			PhoneNumber: phoneNumber,
			AddedAt:     getNowMicro(),
		}
		searches = append(searches, latestSearches)
		config.ApiConfig.LatestSearchesMap[phoneNumber] = latestSearches
	}
	sortLatestSearches(searches)
	searchesUpdate := searches[:getMax(len(searches))]
	config.ApiConfig.LatestSearchesMap = make(map[string]models.LatestSearches, 0)
	for _, latestSearches := range searchesUpdate {
		config.ApiConfig.LatestSearchesMap[latestSearches.PhoneNumber] = models.LatestSearches{
			PhoneNumber: latestSearches.PhoneNumber,
			AddedAt:     latestSearches.AddedAt,
		}
	}
}

func getMax(len int) int {
	if len > 10 {
		return 10
	}
	return len
}

func sortLatestSearches(searches []models.LatestSearches) {
	sort.Slice(searches, func(i, j int) bool {
		return searches[i].AddedAt > searches[j].AddedAt
	})
}

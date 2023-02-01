package service

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SetCronForIndexMap() {
	numbers := commentsUpdatedToday()
	log.Error(numbers)
	cronIndexNow()
}

func cronIndexNow() {
	log.Infof("Setting cronjob for: [Notify IndexNow]")
	c := cron.New()
	_, err := c.AddFunc("@midnight", func() {
		baseUrl := "https://www.example.com/"
		bingUrl := "https://www.bing.com/indexnow?url="
		key := "&key=xxxxxxxxxxxxxxxxxxxxx"
		numbers := commentsUpdatedToday()
		for _, number := range numbers {
			url := bingUrl + baseUrl + number.PhoneNumber + key
			_, err := http.Get(url)
			if err != nil {
				log.Errorf("indexNow http call error: %s", err.Error())
			} else {
				log.Infof("indexNow http call OK: %s", number.PhoneNumber)
			}
		}
	})
	if err != nil {
		log.Panic("cron set error: {}", err.Error())
	}
	c.Start()
}

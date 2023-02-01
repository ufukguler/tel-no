package latest_bar

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func SetLatestCron() {
	initLatestCron()
}

func initLatestCron() {
	log.Infof("Setting cronjob for: [Latest Update]")
	c := cron.New()
	_, err := c.AddFunc("@every 5m", func() {
		InitLatestComments()
		InitLatestSearches()
	})
	if err != nil {
		log.Panic("cron set error: {}", err.Error())
	}
	c.Start()
}

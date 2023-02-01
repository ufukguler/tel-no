package sitemap

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"telno/config"
)

// SetCronForSitemap regenerates the sitemap each hour
func SetCronForSitemap() {
	cronSiteMapUpdate()
	cronSiteMapNotifyDaily()
}

// cronSiteMapUpdate regenerates sitemap every hour
func cronSiteMapUpdate() {
	log.Infof("Setting cronjob for: [Sitemap Update]")
	c := cron.New()
	_, err := c.AddFunc("@midnight", func() {
		config.ApiConfig.Sitemaps = GenerateNumbersSitemap()
		config.ApiConfig.Sitemap = GenerateBaseSitemap()
		config.ApiConfig.SitemapMain = GenerateMainSitemap()
	})
	if err != nil {
		log.Panic("cron set error: {}", err.Error())
	}
	c.Start()
}

// cronSiteMapNotifyDaily notifies search engines at midnight
func cronSiteMapNotifyDaily() {
	log.Infof("Setting cronjob for: [Notify Search Engine]")
	c := cron.New()
	_, err := c.AddFunc("@midnight", func() {
		notifySearchEngines(config.ApiConfig.Sitemap)
	})
	if err != nil {
		log.Panic("cron set error: {}", err.Error())
	}
	c.Start()
}

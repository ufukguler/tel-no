package config

import (
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"os"
	"telno/models"
)

type myConfig struct {
	SiteBaseUrl       string
	SiteMapBaseUrl    string
	MongoHost         string
	MongoPort         string
	MongoUser         string
	MongoPass         string
	Sitemap           *stm.Sitemap
	Sitemaps          []*stm.Sitemap
	SitemapMain       []byte
	CaptchaSecret     string
	EnableSSL         string
	AdminPass         string
	LatestCommentsMap map[string]models.LatestComments
	LatestSearchesMap map[string]models.LatestSearches
	CurseWords        map[string]bool
}

var ApiConfig *myConfig

func LoadConfig() {
	if err := os.Setenv("TZ", "Europe/Istanbul"); err != nil {
		panic(err)
	}
	ApiConfig = new(myConfig)
	ApiConfig.MongoHost = GetEnv("MONGODB_HOST")
	ApiConfig.MongoPort = GetEnv("MONGODB_PORT")
	ApiConfig.MongoUser = GetEnv("MONGODB_USER")
	ApiConfig.MongoPass = GetEnv("MONGODB_ROOT_PASSWORD")
	ApiConfig.EnableSSL = GetEnv("ENABLE_SSL")
	ApiConfig.SiteBaseUrl = GetEnv("SITE_BASE_URL")
	ApiConfig.SiteMapBaseUrl = GetEnv("SITEMAP_BASE_URL")
	ApiConfig.CaptchaSecret = GetEnv("RECAPTCHA_SECRET")
	ApiConfig.LatestCommentsMap = make(map[string]models.LatestComments, 0)
	ApiConfig.LatestSearchesMap = make(map[string]models.LatestSearches, 0)
	ApiConfig.AdminPass = GetEnv("ADMIN_PASS")
}

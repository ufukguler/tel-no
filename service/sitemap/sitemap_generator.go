package sitemap

import (
	"encoding/xml"
	"fmt"
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"telno/config"
	"telno/model_entity"
	"telno/models"
	"telno/service"
	"time"
)

const sitemapLimit = 500

func GenerateMainSitemap() []byte {
	sitemaps := make([]models.MainXMLElement, 0)
	baseUrl := config.ApiConfig.SiteMapBaseUrl

	sitemaps = append(sitemaps, models.MainXMLElement{
		Loc:     baseUrl + "/api/sitemap/pages.xml",
		Lastmod: time.Now().Format(time.RFC3339),
	})
	for i := 0; i < len(config.ApiConfig.Sitemaps); i++ {
		sitemaps = append(sitemaps, models.MainXMLElement{
			Loc:     baseUrl + "/api/sitemap/numbers/" + strconv.Itoa(i+1) + ".xml",
			Lastmod: time.Now().Format(time.RFC3339),
		})
	}
	xmlData := models.MainXML{
		XMLName: xml.Name{
			Space: "1.0",
			Local: "UTF-8",
		},
		Xmlns:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		Sitemap: sitemaps,
	}
	file, _ := xml.MarshalIndent(xmlData, "", " ")
	return []byte(xml.Header + string(file))
}

func GenerateBaseSitemap() *stm.Sitemap {
	sm := stm.NewSitemap(1)
	sm.Create()
	sm.SetDefaultHost(config.ApiConfig.SiteMapBaseUrl)
	sm.SetSitemapsPath("/api/sitemap")

	addToSitemap("", sm)

	return sm
}

func GenerateNumbersSitemap() []*stm.Sitemap {
	crawledUrls, err := service.FindAllForSitemap()
	if err != nil {
		log.Errorf("Sitemap Generator Error: %s", err.Error())
		panic(err)
	}
	sitemaps := make([]*stm.Sitemap, 0)

	loop := int(math.Ceil(float64(len(crawledUrls)) / sitemapLimit))
	for i := 0; i < loop; i++ {
		min := i * sitemapLimit
		max := (i + 1) * sitemapLimit
		if max > len(crawledUrls) {
			max = len(crawledUrls)
		}
		subList := crawledUrls[min:max]
		sm := stm.NewSitemap(1)
		generateSitemap(sm, subList)
		sitemaps = append(sitemaps, sm)
	}
	return sitemaps
}

func generateSitemap(sm *stm.Sitemap, crawledUrls []model_entity.CrawledUrl) {
	if len(crawledUrls) > sitemapLimit {
		panic(fmt.Sprintf("Sitemap array'i %d item'ı geçmemeli!", sitemapLimit))
	}

	baseUrl := config.ApiConfig.SiteBaseUrl
	sm.Create()
	sm.SetDefaultHost(baseUrl)
	sm.SetSitemapsHost(baseUrl)

	for _, v := range crawledUrls {
		sm.Add(stm.URL{
			{"loc", getLocation(baseUrl, v.PhoneNumber)},
			{"lastmod", getUpdateDate(v)},
		})
	}
}

func addToSitemap(path string, sm *stm.Sitemap) {
	baseUrl := config.ApiConfig.SiteBaseUrl
	sm.Add(stm.URL{
		{"loc", baseUrl + path},
		{"lastmod", time.Now().Format(time.RFC3339)},
	})
}

func getLocation(baseUrl, phoneNumber string) string {
	return baseUrl + "/" + phoneNumber
}

func getUpdateDate(v model_entity.CrawledUrl) string {
	if v.UpdatedDate.Year() != 1 {
		return v.UpdatedDate.Format(time.RFC3339)
	}
	return v.CreatedDate.Format(time.RFC3339)
}

func notifySearchEngines(sm *stm.Sitemap) {
	urls := []string{
		"https://api.example.com/api/sitemap/main.xml",
	}
	sm.PingSearchEngines(urls...)
}

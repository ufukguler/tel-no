package models

import "encoding/xml"

type MainXML struct {
	XMLName xml.Name         `xml:"sitemapindex"`
	Xmlns   string           `xml:"xmlns,attr"`
	Sitemap []MainXMLElement `xml:"sitemap"`
}
type MainXMLElement struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

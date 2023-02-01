package models

type Latest struct {
	LatestComments []LatestComments `json:"comments"`
	LatestSearches []LatestSearches `json:"searches"`
}

type LatestComments struct {
	PhoneNumber string `json:"phoneNumber"`
	Comment     string `json:"comment"`
	AddedAt     int64  `json:"addedAt"`
}

type LatestSearches struct {
	PhoneNumber string `json:"phoneNumber"`
	AddedAt     int64  `json:"addedAt"`
}

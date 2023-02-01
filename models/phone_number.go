package models

import (
	"telno/model_entity"
	"time"
)

type CrawledUrlResponseDTO struct {
	PhoneNumber     string                 `json:"phoneNumber"`
	Reading         string                 `json:"reading"`
	Comments        []model_entity.Comment `json:"comments"`
	Operator        string                 `json:"operator"`
	PhoneType       string                 `json:"phoneType"`
	PhoneRegion     string                 `json:"phoneRegion"`
	ReviewCount     int                    `json:"reviewCount"`
	SimilarNumbers  []SimilarNumbersDTO    `json:"similarNumbers,omitempty"`
	AggregateRating AggregateRatingDTO     `json:"aggregateRating"`
}

type SimilarNumbersDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Comment     string `bson:"comment" json:"comment,omitempty"`
	CommentType string `bson:"commentType" json:"commentType,omitempty"`
}

type FindByPhoneNumberRequestDTO struct {
	PhoneNumber string `validate:"required" query:"phoneNumber"`
}

type PhoneNumberImageRequestDTO struct {
	PhoneNumber string `validate:"required" query:"phoneNumber"`
}

type AddCommentRequestDTO struct {
	PhoneNumber string `validate:"required" json:"phoneNumber"`
	Comment     string `validate:"required" json:"comment"`
	CommentType string `validate:"required" json:"commentType"`
	ResponseKey string `validate:"required" json:"responseKey"`
}

type FindByCommentIdRequestDTO struct {
	ID string `validate:"required" query:"id"`
}

type UpdateByCommentIdRequestDTO struct {
	ID          string `validate:"required" json:"id"`
	Comment     string `validate:"required" json:"comment"`
	CommentType string `validate:"required" json:"commentType"`
}

type QuickUpdateByCommentIdRequestDTO struct {
	ID string `validate:"required" json:"id"`
}

type PageableRequestDTO struct {
	Limit       string `validate:"required" query:"limit"`
	Page        string `validate:"required" query:"page"`
	PhoneNumber string `query:"phoneNumber"`
}

type PageableResponseDTO struct {
	Data     interface{} `json:"data"`
	Pageable interface{} `json:"pageable"`
}

type AggregateRatingDTO struct {
	PhoneNumber   string    `json:"phoneNumber"`
	RatingValue   string    `json:"ratingValue"`
	ReviewCount   string    `json:"reviewCount"`
	BestRating    string    `json:"bestRating"`
	WorstRating   string    `json:"worstRating"`
	ReviewBody    string    `json:"reviewBody"`
	PositiveNotes string    `json:"positiveNotes"`
	NegativeNotes string    `json:"negativeNotes"`
	DatePublished time.Time `json:"datePublished"`
}

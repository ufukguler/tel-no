package service

import (
	"fmt"
	"math"
	"telno/model_entity"
	"telno/models"
)

func GenerateAggregateRating(number model_entity.CrawledUrl) models.AggregateRatingDTO {
	reviewCount := 0
	bestRating := "5"
	worstRating := "1"
	reviewBody := ""
	positiveNotes := ""
	negativeNotes := ""

	for _, comment := range number.Comments {
		reviewCount += 1
		if comment.CommentType == "RELIABLE" {
			bestRating = "5"
			positiveNotes = comment.Comment
		} else if comment.CommentType == "DANGEROUS" {
			worstRating = "1"
			negativeNotes = comment.Comment
		}
	}
	if number.Comments != nil && len(number.Comments) > 0 {
		reviewBody = number.Comments[0].Comment
	}

	return models.AggregateRatingDTO{
		PhoneNumber:   "0" + number.PhoneNumber,
		RatingValue:   calculateRating(number.Comments),
		ReviewCount:   fmt.Sprintf("%.1d", reviewCount),
		BestRating:    bestRating,
		WorstRating:   worstRating,
		ReviewBody:    reviewBody,
		PositiveNotes: positiveNotes,
		NegativeNotes: negativeNotes,
		DatePublished: number.CreatedDate,
	}

}

func calculateRating(comments []model_entity.Comment) string {
	if comments != nil && len(comments) > 0 {
		positive := 0
		negative := 0
		for _, comment := range comments {
			if comment.CommentType == "RELIABLE" {
				positive = positive + 1
			} else if comment.CommentType == "DANGEROUS" {
				negative = negative + 1
			}
		}
		total := positive + negative
		if total < 1 {
			total = 1
		}
		ratio := float64(positive / total)
		rating := ratio * 5
		round := math.Round(rating)
		if round < 1 {
			return "1"
		}
		return fmt.Sprintf("%.1f", rating)
	}
	return "2,5"
}

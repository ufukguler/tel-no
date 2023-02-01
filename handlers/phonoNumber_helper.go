package handlers

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"telno/model_entity"
	"telno/models"
	"telno/service"
	"time"
)

func addNewCommentToArray(dto models.AddCommentRequestDTO, crawledUrl model_entity.CrawledUrl) model_entity.CrawledUrl {
	comment := model_entity.Comment{
		Id:          primitive.NewObjectID(),
		Comment:     dto.Comment,
		CommentType: dto.CommentType,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	crawledUrl.Comments = append(crawledUrl.Comments, comment)
	crawledUrl.UpdatedDate = time.Now()
	return crawledUrl
}

func initFirstTimeCrawledUrl(dto models.AddCommentRequestDTO, comments []model_entity.Comment) model_entity.CrawledUrl {
	savedCrawledUrl := model_entity.CrawledUrl{
		Id:          primitive.NewObjectID(),
		PhoneNumber: service.NormalizePhoneNumber(dto.PhoneNumber),
		Comments:    comments,
		Active:      true,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	return savedCrawledUrl
}

func initCommentsForFirstTimeCrawledUrl(dto models.AddCommentRequestDTO) []model_entity.Comment {
	comments := make([]model_entity.Comment, 0)
	comments = append(comments, model_entity.Comment{
		Id:          primitive.NewObjectID(),
		Comment:     dto.Comment,
		CommentType: dto.CommentType,
		Updated:     false,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	})
	return comments
}

func initNotFoundNumber(phoneNumber string) models.CrawledUrlResponseDTO {
	reading := ""
	number := service.NormalizePhoneNumber(phoneNumber)
	if len(number) == 10 {
		reading = service.GetWritingOfPhoneNumber10Digit(number)
	}
	if len(number) == 7 && strings.HasPrefix(number, "444") {
		reading = service.GetWritingOfPhoneNumber7Digit(number)
	}
	region := initPhoneRegion(number)
	phoneType := initPhoneType(number)
	operator := initOperator(number)
	if region != "" {
		// todo
	}

	respArr := initSimilarNumbers(number)
	return models.CrawledUrlResponseDTO{
		PhoneNumber:    number,
		Comments:       make([]model_entity.Comment, 0),
		Reading:        reading,
		Operator:       operator,
		PhoneType:      phoneType,
		PhoneRegion:    region,
		SimilarNumbers: respArr,
		ReviewCount:    1,
		AggregateRating: service.GenerateAggregateRating(model_entity.CrawledUrl{
			PhoneNumber: number,
			Comments:    make([]model_entity.Comment, 0),
			CreatedDate: time.Now(),
		}),
	}
}

func initSimilarNumbers(number string) []models.SimilarNumbersDTO {
	respArr := make([]models.SimilarNumbersDTO, 0)
	similarNumbers, err := service.FindSimilarNumbers(number)
	if err == nil {
		for _, v := range similarNumbers {
			var comment string
			var commentType string
			if len(v.Comments) > 0 {
				comment = v.Comments[0].Comment
				commentType = v.Comments[0].CommentType
			}
			respArr = append(respArr, models.SimilarNumbersDTO{
				PhoneNumber: v.PhoneNumber,
				Comment:     comment,
				CommentType: commentType,
			})
		}
	}
	return respArr
}

func checkCommentType(commentType string) error {
	if commentType == "RELIABLE" || commentType == "DANGEROUS" || commentType == "NEUTRAL" {
		return nil
	}
	return errors.New("invalid comment type")
}

func initOperator(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "444") ||
		strings.HasPrefix(phoneNumber, "501") || strings.HasPrefix(phoneNumber, "505") ||
		strings.HasPrefix(phoneNumber, "506") || strings.HasPrefix(phoneNumber, "507") ||
		strings.HasPrefix(phoneNumber, "551") || strings.HasPrefix(phoneNumber, "552") ||
		strings.HasPrefix(phoneNumber, "553") || strings.HasPrefix(phoneNumber, "554") ||
		strings.HasPrefix(phoneNumber, "555") || strings.HasPrefix(phoneNumber, "559") {
		return "Türk Telekom"
	}
	if strings.HasPrefix(phoneNumber, "516") || strings.HasPrefix(phoneNumber, "530") ||
		strings.HasPrefix(phoneNumber, "531") || strings.HasPrefix(phoneNumber, "532") ||
		strings.HasPrefix(phoneNumber, "533") || strings.HasPrefix(phoneNumber, "534") ||
		strings.HasPrefix(phoneNumber, "535") || strings.HasPrefix(phoneNumber, "536") ||
		strings.HasPrefix(phoneNumber, "537") || strings.HasPrefix(phoneNumber, "538") ||
		strings.HasPrefix(phoneNumber, "539") || strings.HasPrefix(phoneNumber, "561") {
		return "Turkcell"
	}
	if strings.HasPrefix(phoneNumber, "541") ||
		strings.HasPrefix(phoneNumber, "542") || strings.HasPrefix(phoneNumber, "543") ||
		strings.HasPrefix(phoneNumber, "544") || strings.HasPrefix(phoneNumber, "545") ||
		strings.HasPrefix(phoneNumber, "546") || strings.HasPrefix(phoneNumber, "547") ||
		strings.HasPrefix(phoneNumber, "548") || strings.HasPrefix(phoneNumber, "549") {
		return "Vodafone"
	}
	if strings.HasPrefix(phoneNumber, "54285") ||
		strings.HasPrefix(phoneNumber, "54286") || strings.HasPrefix(phoneNumber, "54287") {
		return "KKTC Telsim"
	}
	if strings.HasPrefix(phoneNumber, "53383") ||
		strings.HasPrefix(phoneNumber, "53384") || strings.HasPrefix(phoneNumber, "53385") ||
		strings.HasPrefix(phoneNumber, "53386") || strings.HasPrefix(phoneNumber, "53387") {
		return "KKTC Telsim"
	}
	return ""
}

func initPhoneType(phoneNumber string) string {
	//todo
	return ""
}

func initPhoneRegion(phoneNumber string) string {
	return cities[phoneNumber[0:3]]
}

var cities = map[string]string{
	"322": "Adana",
	"424": "Elazığ",
	"236": "Manisa",
	"416": "Adıyaman",
	"446": "Erzincan",
	"482": "Mardin",
	"272": "Afyon",
	"442": "Erzurum",
	"324": "Mersin",
	"472": "Ağrı",
	"222": "Eskişehir",
	"252": "Muğla",
	"382": "Aksaray",
	"342": "Gaziantep",
	"436": "Muş",
	"358": "Amasya",
	"454": "Giresun",
	"384": "Nevşehir",
	"312": "Ankara",
	"456": "Gümüşhane",
	"388": "Niğde",
	"242": "Antalya",
	"438": "Hakkâri",
	"452": "Ordu",
	"478": "Ardahan",
	"326": "Hatay",
	"328": "Osmaniye",
	"466": "Artvin",
	"476": "Iğdır",
	"464": "Rize",
	"256": "Aydın",
	"246": "Isparta",
	"264": "Sakarya",
	"266": "Balıkesir",
	"216": "İstanbul (Anadolu)",
	"212": "İstanbul (Avrupa)",
	"362": "Samsun",
	"378": "Bartın",
	"484": "Siirt",
	"488": "Batman",
	"232": "İzmir",
	"368": "Sinop",
	"458": "Bayburt",
	"344": "Kahramanmaraş",
	"346": "Sivas",
	"228": "Bilecik",
	"370": "Karabük",
	"414": "Şanlıurfa",
	"426": "Bingöl",
	"338": "Karaman",
	"486": "Şırnak",
	"434": "Bitlis",
	"474": "Kars",
	"282": "Tekirdağ",
	"374": "Bolu",
	"366": "Kastamonu",
	"356": "Tokat",
	"248": "Burdur",
	"352": "Kayseri",
	"462": "Trabzon",
	"224": "Bursa",
	"318": "Kırıkkale",
	"428": "Tunceli",
	"286": "Çanakkale",
	"288": "Kırklareli",
	"276": "Uşak",
	"376": "Çankırı",
	"386": "Kırşehir",
	"432": "Van",
	"364": "Çorum",
	"348": "Kilis",
	"226": "Yalova",
	"258": "Denizli",
	"262": "Kocaeli",
	"354": "Yozgat",
	"412": "Diyarbakır",
	"332": "Konya",
	"372": "Zonguldak",
	"380": "Düzce",
	"274": "Kütahya",
	"284": "Edirne",
	"422": "Malatya",
}

package image_service

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/draw"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"telno/service"
)

func TextOnImg(number string) (image.Image, error) {
	value, err := getGetValue(number)
	if err != nil {
		log.Errorf("getGetValue error: %s", err.Error())
		return nil, err
	}
	var imgPath string
	if value == 0 {
		imgPath = "images/gray.png"
	} else if value > 0 {
		imgPath = "images/green.png"
	} else {
		imgPath = "images/red.png"
	}

	bgImage, err := gg.LoadImage(imgPath)
	if err != nil {
		return nil, err
	}
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Errorf("get font error: %s", err.Error())
		return nil, err
	}
	face := truetype.NewFace(font, &truetype.Options{Size: 80})

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)
	dc.SetFontFace(face)
	maxWidth := float64(imgWidth) - 20.0

	dc.SetColor(color.White)
	if len(number) == 10 {
		dc.DrawStringWrapped(formatPhoneNumber(number), 0, float64(imgHeight/2-30), 0, 0, maxWidth, 1.5, gg.AlignCenter)
	} else {
		dc.DrawStringWrapped(formatPhoneNumber(number), 0, float64(imgHeight/2-30), 0.10, 0, maxWidth, 1.5, gg.AlignCenter)
	}
	//return resizeImage(dc.Image()), nil
	return dc.Image(), nil
}

func resizeImage(src image.Image) *image.RGBA {
	srcRect := src.Bounds()
	dstRect := image.Rectangle{
		Min: image.Point{},
		Max: image.Point{500, 376},
	}
	dst := image.NewRGBA(dstRect)
	draw.CatmullRom.Scale(dst, dstRect, src, srcRect, draw.Over, nil)
	return dst
}

func formatPhoneNumber(number string) string {
	if len(number) == 10 {
		return "+90 " + number[:3] + " " + number[3:6] + " " + number[6:8] + " " + number[8:10]
	}
	if len(number) == 7 {
		return "+90 " + number[:3] + " " + number[3:5] + " " + number[5:7]
	}
	return ""
}

func getGetValue(number string) (int, error) {
	phoneNumber, err := service.FindByPhoneNumber(number)
	if err != nil {
		return 0, nil
	}
	totalPoint := 0
	for _, v := range phoneNumber.Comments {
		if v.Updated {
			totalPoint = totalPoint + evaluateCommentType(v.CommentType)
		}
	}
	return totalPoint, nil
}

func evaluateCommentType(commentType string) int {
	if commentType == "RELIABLE" {
		return 1
	} else if commentType == "DANGEROUS" {
		return -1
	}
	return 0
}

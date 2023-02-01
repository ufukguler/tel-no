package service

import (
	"strconv"
	"strings"
)

var birler, birler2, onlar, onlar2, yuzler []string

func init() {
	birler = []string{"", "bir", "iki", "üç", "dört", "beş", "altı", "yedi", "sekiz", "dokuz"}
	onlar = []string{"", "on", "yirmi", "otuz", "kırk", "elli", "altmış", "yetmiş", "seksen", "doksan"}
	yuzler = []string{"sıfır", "yüz", "ikiyüz", "üçyüz", "dörtyüz", "beşyüz", "altıyüz", "yediyüz", "sekizyüz", "dokuzyüz"}

	birler2 = []string{"sıfır", "bir", "iki", "üç", "dört", "beş", "altı", "yedi", "sekiz", "dokuz"}
	onlar2 = []string{"sıfır", "on", "yirmi", "otuz", "kırk", "elli", "altmış", "yetmiş", "seksen", "doksan"}
}

func GetWritingOfPhoneNumber10Digit(number string) string {
	if len(number) != 10 {
		return ""
	}
	n := []rune(number)

	uc1 := yuzler[toInt(string(n[0]))] + " " + onlar[toInt(string(n[1]))] + " " + birler[toInt(string(n[2]))]
	uc2 := yuzler[toInt(string(n[3]))] + " " + onlar[toInt(string(n[4]))] + " " + birler[toInt(string(n[5]))]

	iki1 := initIkilik(string(n[6]), string(n[7]), toInt(string(n[6])), toInt(string(n[7])))
	iki2 := initIkilik(string(n[8]), string(n[9]), toInt(string(n[8])), toInt(string(n[9])))

	result := strings.ReplaceAll(uc1+" "+uc2+" "+iki1+" "+iki2, "  ", " ")
	result = strings.ReplaceAll(result, "  ", " ")
	return strings.Trim(result, " ")
}

func GetWritingOfPhoneNumber7Digit(number string) string {
	if len(number) != 7 {
		return ""
	}
	n := []rune(number)

	uc1 := yuzler[toInt(string(n[0]))] + " " + onlar[toInt(string(n[1]))] + " " + birler[toInt(string(n[2]))]

	iki1 := initIkilik(string(n[3]), string(n[4]), toInt(string(n[3])), toInt(string(n[4])))
	iki2 := initIkilik(string(n[5]), string(n[6]), toInt(string(n[5])), toInt(string(n[6])))

	result := strings.ReplaceAll(uc1+" "+iki1+" "+iki2, "  ", " ")
	result = strings.ReplaceAll(result, "  ", " ")
	return strings.Trim(result, " ")
}

func toInt(rakam string) int32 {
	i, _ := strconv.ParseInt(rakam, 32, 10)
	return int32(i)
}

func initIkilik(first string, second string, firstIndex int32, secondIndex int32) string {
	first = onlar2[firstIndex]
	second = birler2[secondIndex]
	if firstIndex == 0 && secondIndex == 0 {
		return "çift sıfır"
	}
	if firstIndex == 0 {
		return first + " " + second
	}
	if firstIndex != 0 && secondIndex == 0 {
		return first
	}
	return first + " " + second
}

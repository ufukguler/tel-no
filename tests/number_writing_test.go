package tests

import (
	"telno/service"
	"testing"
)

var numberTest10Digit = []struct {
	got  string
	want string
}{
	{"5428279139", "beşyüz kırk iki sekizyüz yirmi yedi doksan bir otuz dokuz"},
	{"2126220628", "ikiyüz on iki altıyüz yirmi iki sıfır altı yirmi sekiz"},
	{"8504840375", "sekizyüz elli dörtyüz seksen dört sıfır üç yetmiş beş"},
	{"2162370728", "ikiyüz on altı ikiyüz otuz yedi sıfır yedi yirmi sekiz"},
	{"3124040350", "üçyüz on iki dörtyüz dört sıfır üç elli"},
	{"3124040305", "üçyüz on iki dörtyüz dört sıfır üç sıfır beş"},
	{"5053180600", "beşyüz beş üçyüz on sekiz sıfır altı çift sıfır"},
	{"5053000000", "beşyüz beş üçyüz çift sıfır çift sıfır"},
	{"5531122773", "beşyüz elli üç yüz on iki yirmi yedi yetmiş üç"},
}

func TestGetNumberWriting10Digit(t *testing.T) {
	for _, test := range numberTest10Digit {
		testResult := service.GetWritingOfPhoneNumber10Digit(test.got)
		if testResult != test.want {
			t.Errorf("\nwant: '%s' -- %s\n got: '%s'", test.want, test.got, testResult)
		}
	}

}

var numberTest7Digit = []struct {
	got  string
	want string
}{
	{"4444423", "dörtyüz kırk dört kırk dört yirmi üç"},
	{"4440011", "dörtyüz kırk dört çift sıfır on bir"},
	{"4440333", "dörtyüz kırk dört sıfır üç otuz üç"},
	{"4440000", "dörtyüz kırk dört çift sıfır çift sıfır"},
}

func TestGetNumberWriting7DIgit(t *testing.T) {
	for _, test := range numberTest7Digit {
		testResult := service.GetWritingOfPhoneNumber7Digit(test.got)
		if testResult != test.want {
			t.Errorf("\nwant: '%s' -- %s\n got: '%s'", test.want, test.got, testResult)
		}
	}

}

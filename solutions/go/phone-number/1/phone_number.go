package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ".", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+1", "")
	// check for leading 1
	// if present, remove first occurence
	if string(phoneNumber[0]) == "1" {
		phoneNumber = strings.Replace(phoneNumber, "1", "", 1)
	}

	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	areaCodeFirst := string(phoneNumber[0])
	exchangeCodeFirst := string(phoneNumber[3])
	if exchangeCodeFirst == "0" || exchangeCodeFirst == "1" {
		return "", errors.New("Exchange code 0 or 1 invalid")
	}
	if areaCodeFirst == "0" || areaCodeFirst == "1" {
		return "", errors.New("Area code 0 or 1 invalid")
	}
	if len(phoneNumber) != 10 ||
		strings.ContainsAny(phoneNumber,
			"abcdefghijklmnopqrstuvwxyz@:!") {
		return "", errors.New("an error occured")
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", errors.New("Number not parsed")
	}
	areaCode := parsedNumber[0:3]
	return areaCode, nil	
}

func SplitNumberInParts(phoneNumber string) (string, string, string) {
	parsedNumber, _ := Number(phoneNumber)
	areaCode := parsedNumber[0:3] // 613
	exchangeCode := parsedNumber[3:6] // 995
	subscriberNumber := parsedNumber[6:10] // 0253

	return areaCode, exchangeCode, subscriberNumber
}

func Format(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", errors.New("Number not parsed")
	}
	fmt.Printf("Parsing number %s\n", parsedNumber)
	areaCode, exchangeCode, subscriberNumber := SplitNumberInParts(parsedNumber)
	fmt.Printf("split numbers are %s %s %s\n",
		areaCode, exchangeCode, subscriberNumber)
	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberNumber), nil
}

// func main() {
// 	Format("613.995.0253")
// }

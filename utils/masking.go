package utils

import (
	"fmt"
	jsonmask "github.com/teambition/json-mask-go"
	"strings"
)

const (
	MASKING_FROM_LEFT   = "LEFT"
	MASKING_FROM_RIGHT  = "RIGHT"
	MASKING_FROM_CENTER = "CENTER"
	MASKING_FULL        = "FULL"
	MASKING_ASTERISK    = "*"
)

func MaskingJson(logID string, jsonParam interface{}, asterisk []map[string]int) string {

	strJson := JsonToString(logID, jsonParam)

	result := stringToAsterisk(strJson, asterisk)

	return result
}

func MaskingXml(logID string, jsonParam interface{}, asterisk []map[string]int) string {

	strJson := XmlToString(logID, jsonParam)

	result := stringToAsterisk(strJson, asterisk)

	return result
}

func stringToAsterisk(strJson string, asterisk []map[string]int) string {

	result := strJson
	for _, data := range asterisk {
		for key, value := range data {

			strKey := strings.Split(key, "-")
			param1 := strKey[0]

			strIndex := strings.Index(result, param1)
			fmt.Println("strIndex ",strIndex)
			fmt.Println("param1 ",param1)
			from := getIndexFromStrJson(strIndex + len(param1))
			to := from + value
			fmt.Println("from ",from)
			fmt.Println("to ",to)
			strToAsterisk := result[from:to]
			fmt.Println("strToAsterisk ",strToAsterisk)
			result = toAsterisk(result, strToAsterisk, strKey[1], value)
		}
	}

	return result
}

func getIndexFromStrJson(strIndex int) int {

	return strIndex + 4
}

func toAsterisk(strJson string, strToAsterisk string, maskingFrom string, maskingLen int) string {

	asterisk := strings.Repeat(MASKING_ASTERISK, maskingLen)
	result := strings.ReplaceAll(strJson, strToAsterisk, asterisk)

	if maskingLen > len(strToAsterisk) {
		return result
	}

	switch maskingFrom {
	case MASKING_FROM_LEFT:
		result = fromLeft(strJson, strToAsterisk, maskingLen)
	case MASKING_FROM_RIGHT:
		result = fromRight(strJson, strToAsterisk, maskingLen)
	case MASKING_FROM_CENTER:
		result = fromCenter(strJson, strToAsterisk, maskingLen)
	case MASKING_FULL:
		return result
	default:
		return result
	}

	return result
}

func fromRight(strJson string, strToAsterisk string, maskingLen int) string {

	fmt.Println("strToAsterisk ", strToAsterisk)
	from := len(strToAsterisk) - maskingLen
	to := len(strToAsterisk)
	newString := strToAsterisk[from:to]
	fmt.Println("from ", from)
	fmt.Println("to ", to)
	fmt.Println("newString ", newString)
	return strToAsterisk[0:len(strToAsterisk)] + strings.Repeat(MASKING_ASTERISK, len(newString))
}

func fromLeft(strJson string, strToAsterisk string, maskingLen int) string {

	from := len(strToAsterisk)
	to := maskingLen
	newString := strToAsterisk[from:to]

	return strings.ReplaceAll(strJson, newString, MASKING_ASTERISK)
}

func fromCenter(strJson string, strToAsterisk string, n int) string {

	from := len(strToAsterisk) + n
	to := len(strToAsterisk) - n
	newString := strToAsterisk[from:to]

	return strings.ReplaceAll(strJson, newString, MASKING_ASTERISK)
}

func Masdf(param interface{}, str string) string {

	result, _ := jsonmask.Mask(param, str)
	s := fmt.Sprintf("%v", result)
	return s
}

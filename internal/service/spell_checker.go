package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"note/internal/types"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	postRequestURL = "https://speller.yandex.net/services/spellservice.json/checkText?text="
)

// Отправка запроса в спеллер и возврат исправленного текста
func useSpellChecker(text string) (string, error) {
	logrus.Info("Sending text to speller: ", text)
	urlText := strings.ReplaceAll(text, " ", "+")
	response, err := http.Post(postRequestURL+urlText, "text/plain", nil)
	if err != nil {
		logrus.Error("error sending request: ", err)
		return "", fmt.Errorf("something went wrong")
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		logrus.Error("error reading from response body: ", err)
		return "", fmt.Errorf("something went wrong")
	}

	var checkResponse []types.CheckText
	err = json.Unmarshal(body, &checkResponse)
	if err != nil {
		logrus.Error("error while unmarshalling: ", err)
		return "", fmt.Errorf("something went wrong")
	}

	return getSpelledText(checkResponse, text), nil
}

// Исправление исходного текста
func getSpelledText(checkResponse []types.CheckText, inputString string) string {
	inputRune := []rune(inputString)
	var result []rune

	var index int
	for _, word := range checkResponse {
		if index >= len(inputString) {
			break
		}
		subString := inputRune[index:word.Pos]
		result = append(result, subString...)
		result = append(result, []rune(word.S[0])...)
		index = word.Pos + word.Len
	}

	if index < len(inputRune) {
		result = append(result, inputRune[index:]...)
	}
	return string(result)
}

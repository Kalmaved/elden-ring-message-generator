package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)


// Each category is a list of words/phrases
type mpd map[string][]string


func main() {

	var messagePartsData mpd = getMessageParts("messagePartsData.json")

	result := getRandomMessage(messagePartsData)
	fmt.Println(result)

}


func getMessageParts(path_to_file string) mpd {

	// opening json and reading from it
	jsonFile, err := os.Open(path_to_file)
	if err != nil { fmt.Println(err) }
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var messagePartsData map[string][]string
	json.Unmarshal([]byte(byteValue), &messagePartsData)
	
	return messagePartsData

}


func getMessageFromTemplate(template string, word string) string {
	return strings.ReplaceAll(template, "****", word)
}


func getRandomMessage(messagePartsData mpd) string {

	rand.Seed(time.Now().UnixNano())

	//var isSimpleFormat bool

	// Get random message format
	//format := formats[rand.Intn(len(formats))]

	templates := messagePartsData["TEMPLATES"]
	conjunctions := messagePartsData["CONJUNCTIONS"]
	categories := make([]string, 0, len(messagePartsData) - 2)
	for key := range messagePartsData {
		if key != "TEMPLATES" && key != "CONJUNCTIONS" {
			categories = append(categories, key)
		}
	}

	var message string

	// random message format; there are two — either 1 template or 2 templates with conjunction
	// adding 1 just for nicer representation
	switch 1 + rand.Int() % 2 {

		case 1:
			randomTemplate := templates[rand.Intn(len(templates))]
			randomCategory := categories[rand.Intn(len(categories))]
			randomWord := messagePartsData[randomCategory][rand.Intn(len(messagePartsData[randomCategory]))]

			message = getMessageFromTemplate(randomTemplate, randomWord)

		case 2:
			randomTemplateFirst := templates[rand.Intn(len(templates))]
			randomCategoryFirst := categories[rand.Intn(len(categories))]
			randomWordFirst := messagePartsData[randomCategoryFirst][rand.Intn(len(messagePartsData[randomCategoryFirst]))]

			randomTemplateSecond := templates[rand.Intn(len(templates))]
			randomCategorySecond := categories[rand.Intn(len(categories))]
			randomWordSecond := messagePartsData[randomCategorySecond][rand.Intn(len(messagePartsData[randomCategorySecond]))]

			randomConjunction := conjunctions[rand.Intn(len(conjunctions))]

			message = fmt.Sprintf("%s\n%s\n%s", getMessageFromTemplate(randomTemplateFirst, randomWordFirst), randomConjunction, getMessageFromTemplate(randomTemplateSecond, randomWordSecond))
	}

	return message
}

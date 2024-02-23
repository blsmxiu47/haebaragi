package main

import (
	"log"
	"net/http"
)

var testWords = []Word{
	{
		ID:                  1,
		SourceText:          "Sunflower",
		TranslationText:     "Sonnenblume",
		SourceLanguage:      "English",
		TranslationLanguage: "Deutsch",
		Category:            "",
		PartOfSpeech:        "noun",
		Gender:              "f",
		Examples:            []string{"Die Sonnenblume ist groß."},
		Notes:               "",
	},
	{
		ID:                  2,
		SourceText:          "German",
		TranslationText:     "Deutsch",
		SourceLanguage:      "English",
		TranslationLanguage: "Deutsch",
		Category:            "",
		PartOfSpeech:        "noun",
		Gender:              "n",
		Examples:            []string{"Ich spreche kein Deutsch."},
		Notes:               "",
	},
}

var testExpressions = []Expression{
	{
		ID:                  1,
		SourceText:          "to be on the ball",
		TranslationText:     "auf Draht sein",
		SourceLanguage:      "English",
		TranslationLanguage: "Deutsch",
		Category:            "",
		Notes:               "",
	},
}

var testLanguages = []Language{
	{
		ID:   1,
		Name: "English",
		Code: "EN",
	},
	{
		ID:   2,
		Name: "Deutsch",
		Code: "DE",
	},
}

var testCategories = []Category{
	{
		ID:   1,
		Name: "Greetings",
	},
	{
		ID:   2,
		Name: "Food",
	},
	{
		ID:   3,
		Name: "People",
	},
}

var testGenders = []Gender{
	{
		ID:           1,
		Name:         "neuter",
		Abbreviation: "n",
		Languages:    []Language{testLanguages[1]},
	},
	{
		ID:           2,
		Name:         "masculine",
		Abbreviation: "m",
		Languages:    []Language{testLanguages[1]},
	},
	{
		ID:           3,
		Name:         "feminine",
		Abbreviation: "f",
		Languages:    []Language{testLanguages[1]},
	},
}

var testPartsOfSpeech = []PartOfSpeech{
	{
		ID:           1,
		Name:         "noun",
		Abbreviation: "n",
	},
	{
		ID:           2,
		Name:         "verb",
		Abbreviation: "v",
	},
	{
		ID:           3,
		Name:         "adjective",
		Abbreviation: "adj",
	},
	{
		ID:           4,
		Name:         "adverb",
		Abbreviation: "adv",
	},
	{
		ID:           5,
		Name:         "preposition",
		Abbreviation: "prep",
	}, // TODO: will add all when we have actual DB set up
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\n해바라기에 환영합니다!\n\n"))

	w.Write([]byte("Endpoints:\n\n"))
	w.Write([]byte("GET /word              POST /word\n"))
	// w.Write([]byte("GET /expression         POST /expression\n"))
	w.Write([]byte("GET /language	       POST /language\n"))
	w.Write([]byte("GET /part-of-speech    POST /part-of-speech\n"))
	w.Write([]byte("GET /gender            POST /gender\n"))
	w.Write([]byte("GET /category	       POST /category\n"))
	w.Write([]byte("\n"))
}

func handleRequests() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/word", wordHandler)
	http.HandleFunc("/expression", expressionHandler)
	http.HandleFunc("/language", languageHandler)
	http.HandleFunc("/part-of-speech", partOfSpeechHandler)
	http.HandleFunc("/gender", genderHandler)
	http.HandleFunc("/category", categoryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

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
	w.Write([]byte("해바라기에 환영합니다!"))
}

func handleRequests() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/word", addWordHandler)
	http.HandleFunc("/expression", addExpressionHandler)
	http.HandleFunc("/language", addLanguageHandler)
	http.HandleFunc("/category", addCategoryHandler)
	http.HandleFunc("/part-of-speech", addPartOfSpeechHandler)
	http.HandleFunc("/gender", addGenderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("해바라기에 환영합니다!"))
}

func handleRequests() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/word", addWordHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

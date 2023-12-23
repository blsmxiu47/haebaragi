package main;

import (
	"net/http"
)

var testWords = []Word {
	{
		ID: 1,
		SourceText: "Sunflower",
		TranslationText: "Sonnenblume",
		SourceLanguage: "English",
		TranslationLanguage: "Deutsch",
		Category: "",
		PartOfSpeech: "",
		Gender: "f",
		Examples: []string{"Die Sonnenblume ist groß."},
		Notes: "",
	},
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/word", addWordHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("해바라기에 환영합니다!"))
}

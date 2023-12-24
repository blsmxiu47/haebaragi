package main

import (
	"encoding/json"
	"net/http"
)

func addWordHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testWords)
}

func addExpressionHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testExpressions)
}

func addLanguageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testLanguages)
}

func addPartOfSpeechHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testPartsOfSpeech)
}

func addGenderHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testGenders)
}

func addCategoryHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testCategories)
}

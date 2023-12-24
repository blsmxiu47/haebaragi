package main

import (
	"encoding/json"
	"net/http"
)

// Testing
var words []Word
var nextWordID int = 1

func addWordToStore(word Word) error {
	word.ID = nextWordID
	nextWordID++
	words = append(words, word)
	return nil
}

// Testing
var languages []Language
var nextLanguageID int = 1

func addLanguageToStore(language Language) error {
	language.ID = nextWordID
	nextWordID++
	languages = append(languages, language)
	return nil
}

// TODO: obvi these handlers will all do basically the same thing but for different data, so refactor
func wordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testWords)
	case "POST":
		var word Word
		err := json.NewDecoder(r.Body).Decode(&word)
		if err != nil {
			http.Error(w, "Error reading word from request", http.StatusBadRequest)
			return
		}

		// TODO: Add validation for word fields

		err = addWordToStore(word)
		if err != nil {
			http.Error(w, "Error saving word", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(word)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func expressionHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testExpressions)
}

// TODO: obvi these handlers will all do basically the same thing but for different data, so refactor
func languageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testLanguages)
	case "POST":
		var language Language
		err := json.NewDecoder(r.Body).Decode(&language)
		if err != nil {
			http.Error(w, "Error reading language from request", http.StatusBadRequest)
			return
		}

		// TODO: Add validation for language fields

		err = addLanguageToStore(language)
		if err != nil {
			http.Error(w, "Error saving language", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(language)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func partOfSpeechHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testPartsOfSpeech)
}

func genderHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testGenders)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testCategories)
}

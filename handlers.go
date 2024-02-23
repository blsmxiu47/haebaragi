package main

import (
	"encoding/json"
	"errors"
	"log"
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

func addLanguageToStore(language Language) error {
	language.ID = nextWordID
	nextWordID++
	languages = append(languages, language)
	return nil
}

// Testing
var partsOfSpeech []PartOfSpeech
var nextPartOfSpeechID int = 1

func addPartOfSpeechToStore(partOfSpeech PartOfSpeech) error {
	partOfSpeech.ID = nextPartOfSpeechID
	nextPartOfSpeechID++
	partsOfSpeech = append(partsOfSpeech, partOfSpeech)
	return nil
}

// Testing
var genders []Gender
var nextGenderID int = 1

func addGenderToStore(gender Gender) error {
	gender.ID = nextGenderID
	nextGenderID++
	genders = append(genders, gender)
	return nil
}

// Testing
var categories []Category
var nextCategoryID int = 1

func addCategoryToStore(category Category) error {
	category.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, category)
	return nil
}

// TODO: obvi these handlers will all do basically the same thing but for different data, so refactor
func wordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testWords)
	case "POST":
		var word Word

		err := decodeJSONBody(w, r, &word)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
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

		err := decodeJSONBody(w, r, &language)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
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
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testPartsOfSpeech)
	case "POST":
		var partOfSpeech PartOfSpeech

		err := decodeJSONBody(w, r, &partOfSpeech)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		// TODO: Add validation for partOfSpeech fields

		err = addPartOfSpeechToStore(partOfSpeech)
		if err != nil {
			http.Error(w, "Error saving partOfSpeech", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(partOfSpeech)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func genderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testGenders)
	case "POST":
		var gender Gender

		err := decodeJSONBody(w, r, &gender)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		// TODO: Add validation for gender fields

		err = addGenderToStore(gender)
		if err != nil {
			http.Error(w, "Error saving gender", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(gender)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testCategories)
	case "POST":
		var category Category

		err := decodeJSONBody(w, r, &category)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		// TODO: Add validation for category fields

		err = addCategoryToStore(category)
		if err != nil {
			http.Error(w, "Error saving category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(category)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	json.NewEncoder(w).Encode(testCategories)
}

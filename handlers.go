package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Constants
const MAX_REQUEST_BODY_SIZE_MSG = "1MB"
const MAX_REQUEST_BODY_SIZE_BYTES = 1048576

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
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			msg := "Invalid content-type. Content-Type is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, MAX_REQUEST_BODY_SIZE_BYTES)

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var word Word
		err := dec.Decode(&word)
		print(r.Body)
		if err != nil {
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError

			switch {
			case errors.As(err, &syntaxError):
				msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
				http.Error(w, msg, http.StatusBadRequest)

			case errors.Is(err, io.ErrUnexpectedEOF):
				msg := "Request body contains badly-formed JSON"
				http.Error(w, msg, http.StatusBadRequest)

			case errors.As(err, &unmarshalTypeError):
				msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
				http.Error(w, msg, http.StatusBadRequest)

			case strings.HasPrefix(err.Error(), "json: unknown field "):
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
				msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
				http.Error(w, msg, http.StatusBadRequest)

			case errors.Is(err, io.EOF):
				msg := "Request body must not be empty"
				http.Error(w, msg, http.StatusBadRequest)

			case err.Error() == "http: request body too large":
				msg := fmt.Sprintf("Request body must not be larger than %s", MAX_REQUEST_BODY_SIZE_MSG)
				http.Error(w, msg, http.StatusRequestEntityTooLarge)

			default:
				log.Print(err.Error())
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

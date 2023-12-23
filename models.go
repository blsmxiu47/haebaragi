package main

type Word struct {
	ID                  int     
	SourceText          string   `json:"id"`
	TranslationText     string   `json:"translationText"`
	SourceLanguage      string   `json:"sourceLanguage"`
	TranslationLanguage string   `json:"translationLanguage"`
	PartOfSpeech        string   `json:"partOfSpeech"`
	Gender              string   `json:"gender"`
	Category            string   `json:"category"`
	Examples            []string `json:"examples"`
	Notes               string   `json:"notes"`
}

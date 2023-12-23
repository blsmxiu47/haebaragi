package main

type Word struct {
	ID                  int
	SourceText          string
	TranslationText     string
	SourceLanguage      string
	TranslationLanguage string
	PartOfSpeech        string
	Gender              string
	Category            string
	Examples            []string
	Notes               string
}

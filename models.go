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

type Expression struct {
	ID                  int
	SourceText          string `json:"id"`
	TranslationText     string `json:"translationText"`
	SourceLanguage      string `json:"sourceLanguage"`
	TranslationLanguage string `json:"translationLanguage"`
	Category            string `json:"category"`
	Notes               string `json:"notes"`
}

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type PartOfSpeech struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Gender struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Abbreviation string     `json:"abbreviation"`
	Languages    []Language `json:"languages"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

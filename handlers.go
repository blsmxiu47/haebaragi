package main

import (
	"encoding/json"
	"net/http"
)

func addWordHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testWords);
}

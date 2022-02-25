package main

import (
    "encoding/json"
    "io"
    "net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Welcome, hubuc\n")
}

type Result struct {
    UniqueID string `json:"uniqueID"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    result, _ := json.Marshal(Result{"12345"})
    io.WriteString(w, string(result))
}

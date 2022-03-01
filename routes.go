package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello, hubuc\n")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "12345\n")
}

type Result struct {
    UniqueID string `json:"uniqueID"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    result, _ := json.Marshal(Result{"12345"})
    io.WriteString(w, string(result))
}

func HandleDebit(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "debit\n")
}

func HandleCredit(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "credit\n")
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "create\n")
}

func HandleGetCardDetails(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "getCardDetails\n")
}

func HandleBlock(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "block\n")
}

func HandleUnblock(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "unblock\n")
}


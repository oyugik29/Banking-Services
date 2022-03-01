package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
	"fmt"
	"io/ioutil"

)

type card struct{
	PAN	string `json:"pan" validate"required"`
}

type allCards []card

var cards = allCards{
	{
		PAN:	"1",
	},
}

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

func HandleGetCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)


}

func HandleDebit(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "debit\n")
}	

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var newCard card
	reqBody, err := ioutil.ReadAll(r.Body)
	if len(reqBody) == 0{
		w.WriteHeader(http.StatusNoContent)
	}
	if err != nil {
		fmt.Fprint(w, "Kindly enter the PAN")
	}
	json.Unmarshal(reqBody, &newCard)

	if len(newCard.PAN) == 0{
		w.WriteHeader(http.StatusNoContent)
	}
	cards = append(cards, newCard)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCard)
}


func HandleCredit(w http.ResponseWriter, r *http.Request) {
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


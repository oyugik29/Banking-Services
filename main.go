package main

import (
    "log"
    "net/http"
)

func main() {
handleRequests()

    log.Fatal(http.ListenAndServe(":9900", nil))
}

func handleRequests(){
    // public views
    http.HandleFunc("/", HandleIndex)

    // private views
    http.HandleFunc("/post", PostOnly(basicAuth(HandlePost)))
    http.HandleFunc("/json", GetOnly(basicAuth(HandleJSON)))
    http.HandleFunc("/debit", PostOnly(basicAuth(HandleDebit)))
    http.HandleFunc("/credit", PostOnly(basicAuth(HandleCredit)))
    http.HandleFunc("/create", PostOnly(basicAuth(HandleCreate)))
    http.HandleFunc("/getCardDetails", PostOnly(basicAuth(HandleGetCardDetails)))
    http.HandleFunc("/block", PostOnly(basicAuth(HandleBlock)))
    http.HandleFunc("/unblock", PostOnly(basicAuth(HandleUnblock)))
    http.HandleFunc("/getCards", GetOnly(basicAuth(HandleGetCards)))


}

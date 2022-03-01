package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    file, err := openLog("bank.log")
    if err != nil{
        log.Fatal(err)
    }
    log.SetOutput(file)
    log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
    log.Println("logFile Created")

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

func openLog(path string) (*os.File, error){
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil{
        return nil, err
    }
    return logFile, nil
}

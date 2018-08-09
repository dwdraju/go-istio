package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"
)

type Version struct {
    Version int
}

func home(w http.ResponseWriter, r *http.Request){
    version := Version{2}
    js, err := json.Marshal(version)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    w.Header().Set("Content-Type","application/json")
    w.Write(js)
}

func main() {
    router := mux.NewRouter()
    port   := 8080
    addr := fmt.Sprintf(":%d", port)
    router.HandleFunc("/", home).Methods("GET")
    fmt.Printf("App is running on http://localhost:%d !\n", port)
    log.Fatal(http.ListenAndServe(addr, router))
}

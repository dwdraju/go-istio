package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"
    "math/rand"
    "time"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request){
    fail, err := strconv.Atoi(r.Header.Get("fail"))
    myrand:= random(1, 500)
    fmt.Println(myrand)
    if fail < myrand {
        http.Error(w, "Oops ! Something went wrong!!", 500)
        return
    }

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    w.Write([]byte("Yay !!! You're the luckiest man on earth!!! "))
}

func random(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}

func main() {
    router := mux.NewRouter()
    port   := 8080
    addr := fmt.Sprintf(":%d", port)
    router.HandleFunc("/", home).Methods("GET")
    fmt.Printf("App is running on http://localhost:%d !\n", port)
    log.Fatal(http.ListenAndServe(addr, router))
}

package main

import (
    "fmt"
    
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {

    http.HandleFunc("/", home)
    
    log.Fatal(http.ListenAndServe(":8081", nil))

}
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./")) // New code
    http.Handle("/", fileServer) // New code


    fmt.Printf("Starting server at port 4000\n")
    if err := http.ListenAndServe(":4000", nil); err != nil {
        log.Fatal(err)
    }
}

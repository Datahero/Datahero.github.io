package main

import (
    "fmt"
    "log"
    "net/http"
)

func artisanHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "artisan.html")
}
func barndoorHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "barndoor.html")
}

func main() {
    fileServer := http.FileServer(http.Dir("./")) // New code
    http.Handle("/", fileServer)
    http.HandleFunc("/artisan", artisanHandler)
    http.HandleFunc("/barndoor", barndoorHandler)


    fmt.Printf("Starting server at port 4000\n")
    if err := http.ListenAndServe(":4000", nil); err != nil {
        log.Fatal(err)
    }
}

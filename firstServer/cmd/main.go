package main

import (
    "fmt"
    "log"
    "net/http"

    myhandler "my-server/internal/http"  
)

func main() {
    fmt.Println("Мой сервер работает на ip 127.0.0.1:8080")

    http.HandleFunc("/hello", myhandler.HelloHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
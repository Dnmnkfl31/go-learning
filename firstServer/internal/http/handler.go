package http

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Привет от моего Go-сервера!")
}
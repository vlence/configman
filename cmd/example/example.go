package main

import (
	"net/http"
    _ "embed"
)

//go:embed templates/base.html
var baseTmpl string

//go:embed templates/index.html
var indexTmpl string

//go:embed templates/config.html
var configTmpl string

func main() {
        http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        })
}

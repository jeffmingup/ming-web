package main

import (
	"fmt"
	"ming"
	"net/http"
)

func main() {
	r := ming.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "url路径:%q\n", r.URL.Path)
	})

	r.Run(":9999")
}

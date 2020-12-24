package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {

}

func (engine *Engine)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w,"url:%s\n",r.URL.Path)
	case "/hello":
		for k ,v :=range r.Header{
			fmt.Fprintf(w,"header[%s]:%s\n",k,v)
		}
	default:
		fmt.Fprint(w,777)
	}
}

func main()  {
	engine:= new(Engine)
	log.Fatal(http.ListenAndServe(":9999",engine))
	
}

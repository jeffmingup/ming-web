package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/hello",helloHandle)
	log.Fatal(http.ListenAndServe(":9999",nil))
	
}
func indexHandle(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"url:%s\n",request.URL.Path)
}
func helloHandle(w http.ResponseWriter,r *http.Request)  {
	for k ,v :=range r.Header{
		fmt.Fprintf(w,"header[%s]:%s\n",k,v)
	}

}
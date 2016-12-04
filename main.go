package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello Golang")
}
func main()  {
	mux:=http.NewServeMux()
	files:=http.FileServer(http.Dir("./public"))
	mux.Handle("/static/",http.StripPrefix("/static",files))
	mux.HandleFunc("/",index)
	server:=&http.Server{
		Addr:"localhost:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}
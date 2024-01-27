package main

import (
	"fmt"
	"log"
	"net/http"
)
	
func helloHandler(w http.ResponseWriter, r *http.Request){// has 2 parameters response and request, r ponts '*' to request 
if r.URL.Path != "/hello" {
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}

if r.Method != "GET" {
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w, "hello!")
}

func formHandler(w, http.ResponseWriter, r *http.Request){
if err := r.ParseForm(); err != nil {
	return
}
fmt.Fprintf(w, "POST request successfull")
name := r.FormValue("name")
address := r.FormValue("address")
fmt.Fprintf(w, "Name = %s\n", name)
fmt.Fprintf(w, "Address = %s\n", address)
}


func main(){
fileServer := htt.FileServer(http.Dir("./static")) //check the static directory
http.Handle("/", fileServer) //handle root route
http.HandleFunc("/form", formHandler)
http.HandleFunc("/hello", helloHandler)

fmt.Println("Starting server at port 8080\n")
if err := http.ListenAndServe(":8080", nil); err != nil {// will create the server
	log.Fatal(err)
}
}
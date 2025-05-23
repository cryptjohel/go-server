package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){

	// if r.Method != "POST"{
	// 	http.Error(w, "Method is not supported", http.StatusNotFound)
	
	if err := r.ParseForm(); err !=nil{
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello my boi")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err !=nil{
		log.Fatal(err)
	}
	
}

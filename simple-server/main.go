package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request){
	if  req.URL.Path != "/hello" && req.Method != "GET"{
		http.Error(res, "404 Not Found", http.StatusNotFound)
		return;
     }
   
   fmt.Fprintf(res, "Hello!")
}


func formHandler(res http.ResponseWriter, req *http.Request){
	error := req.ParseForm()
	if error != nil {
        fmt.Fprintf(res, "ParseForm() error: %v", error)
		return;
	}
	
	fmt.Fprintf(res, "POST request successful.")

	name :=req.FormValue("name")
	address :=req.FormValue("address")

	fmt.Fprintf(res, "Name = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Server is on port 8081 \n")

	error := http.ListenAndServe(":8081", nil)

	if error != nil {
		log.Fatal(error)
	}
}
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	fmt.Println("Hello mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	log.Fatal(http.ListenAndServe(":8080", r))

 
}

func greeter(){
	fmt.Println("Hey there mod users")
}

func  serveHome(w http.ResponseWriter, r *http.Request){
	// Add this to see if the browser is connecting
    fmt.Println("Request received at /")

	w.Write([]byte("<h1>Welcome to the golang series on Yt </h1>"))
}


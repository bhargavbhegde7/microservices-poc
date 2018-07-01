package main

import (
	"fmt"
	"net/http"
	"log"
	//"time"
	"io/ioutil"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetIndexEndpoint(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
		
	}
	req, err := http.NewRequest("GET", "http://localhost:3000/api/auth/verify-token", nil)
	if err!= nil{
		fmt.Fprintf(w, "Token verification failed")
		log.Fatal("Couldn't create the request object for token verification")
	}

	req.Header.Set("x-access-token", r.Header.Get("x-access-token"))
	resp, err := client.Do(req)
	if err!= nil{
		fmt.Fprintf(w, "Token verification failed")
		log.Fatal("Http GET request failed")
	}else{
		body, err := ioutil.ReadAll(resp.Body)
		if err!= nil{
			fmt.Fprintf(w, "Token verification failed")
			log.Fatal("Couldn't parse the response body")
		}
		bodyString := string(body)
		fmt.Fprintf(w, bodyString)
	}
}

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	//router.HandleFunc("/", PostIndexEndpoint).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"x-access-token"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":" + "5000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
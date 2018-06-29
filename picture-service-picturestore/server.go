package main

import (
	"fmt"
	"net/http"
	"log"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
    name string
    email string
}

//GetIndexEndpoint . . .
func GetIndexEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("x-access-token"))
	
	

	fmt.Fprintf(w, "These are not the droids you're looking for")
}

//PostIndexEndpoint . . .
/*func GetIndexEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("x-access-token"))
	
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	//resp, err := client.Get("http://auth-service-picturestore:3000/api/auth/verify-token")
	resp, err := client.Get("http://localhost:3000/api/auth/verify-token")
	
	fmt.Fprintf(w, "Here is all your pictures pleb")
}*/

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	//router.HandleFunc("/", PostIndexEndpoint).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"x-access-token"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":" + "5000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
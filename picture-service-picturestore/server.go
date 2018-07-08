package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"io/ioutil"
	"crypto/md5"
	"io"
	"strconv"
	"text/template"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// upload logic
func PostIndexEndpoint(w http.ResponseWriter, r *http.Request) {
       fmt.Println("method:", r.Method)
       if r.Method == "GET" {
           crutime := time.Now().Unix()
           h := md5.New()
           io.WriteString(h, strconv.FormatInt(crutime, 10))
           token := fmt.Sprintf("%x", h.Sum(nil))

           t, _ := template.ParseFiles("upload.gtpl")
           t.Execute(w, token)
       } else {
           r.ParseMultipartForm(32 << 20)
           file, handler, err := r.FormFile("uploadfile")
           if err != nil {
               fmt.Println(err)
               return
           }
           defer file.Close()
           fmt.Fprintf(w, "%v", handler.Header)
           f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
           if err != nil {
               fmt.Println(err)
               return
           }
           defer f.Close()
           io.Copy(f, file)
       }
}

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
		//fmt.Fprintf(w, bodyString)
		fmt.Println(bodyString)//real response

		//mock response with picture urls.
		b, err := ioutil.ReadFile("format.json") // b has type []byte
		if err != nil {
			log.Fatal(err)
		}
		str := string(b)
		fmt.Fprintf(w, str)
	}
}

/*func GetImagesEndpoint(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("format.json") // b has type []byte
		if err != nil {
			log.Fatal(err)
		}
		str := string(b)
		fmt.Fprintf(w, b)
}*/

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	//router.HandleFunc("/images/{img-id}", GetImagesEndpoint).Methods("GET")	
	router.HandleFunc("/", PostIndexEndpoint).Methods("POST")

    /* --------------------------------- */
    router.PathPrefix("/static").Handler(http.FileServer(http.Dir("assets/")))

	headersOk := handlers.AllowedHeaders([]string{"x-access-token"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":" + "5000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	//log.Fatal(http.ListenAndServe(":" + "5000", nil))
}
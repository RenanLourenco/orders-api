package main

import (
	"log"
	"net/http"
)

func main(){
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to listen the server")
	}
}

func basicHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello, world"))

}
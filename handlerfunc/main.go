package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr: ":8080",
		// http.HandlerFunc(...) looks like a function call but it's a type conversion.
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Dropout Boogie"))
		}),
	}

	log.Println("Listening on http://localhost:8080")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

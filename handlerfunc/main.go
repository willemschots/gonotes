// # Using http.HandlerFunc
//
// This command runs a http handler using http.HandlerFunc. This is a small demo related to my [article on handler func] 👨‍🎤.
//
// There is an executable `main.go` that contains a demo.
//
// # Notes
//
//   - Functions as values.
//   - Function types that call themselves to implement an interface.
//   - Use compiler errors to guide explanation.
//   - http.HandlerFunc and http.HandleFunc have very similar names and can be confused.
//   - http.HandlerFunc(...) looks like a function call but it's a type conversion.
//   - http.HandlerFunc implements http.Handler by calling itself. This can be a bit trippy for newcomers.
//
// [article on handler func]: https://www.willem.dev/articles/http-handler-func/
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

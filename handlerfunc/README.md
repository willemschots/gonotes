# Playing around with http.HandlerFunc

This package contains my personal notes on `http.HandlerFunc`. These notes are the basis of my [article on handler func](https://www.willem.dev/articles/http-handler-func/) 👨‍🎤.

There is an executable `main.go` that contains a demo.

## Notes:

- User needs to know about:
  - Functions as values.
  - Interfaces.
- Use compiler errors to guide explanation.
- `http.HandlerFunc` and `http.HandleFunc` have very similar names and can be confused.
- `http.HandlerFunc(...)` looks like a function call but it's a type conversion.
- `http.HandlerFunc` implements `http.Handler` by calling itself. This can be a bit trippy for newcomers.

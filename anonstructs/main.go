// # Using anonymous structs
//
// This command explores anonymous structs. This is a small demo related to my [article on anonymous structs].
//
// # Notes
//
//   - Anonymous structs are structs defined without a name.
//   - Also called "inline structs".
//   - Limited utility, but can emphasize that a struct is only used in a single place.
//   - Common in the table test pattern.
//   - Can become hard to read when the struct is large.
//
// [article on anonymous structs]: https://www.willem.dev/articles/anonymous-structs/
package main

import "fmt"

func main() {
	d := struct {
		s string
		x int
		y float64
	}{
		s: "hello world",
		x: 101,
		y: 3.14,
	}

	fmt.Println(d)
}

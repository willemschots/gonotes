// # Time in Go
//
// This command explores time in Go. This is a small demo related to my articles on time:
//
//   - [How to parse a time or date in Go]
//   - [time.Time and time.Location explained]
//   - [time.Now and the monotonic clock]
//   - [Comparing a time or date in Go]
//
// # Notes
//
//   - Go uses a so called "reference layout" for formatting and parsing time.Time structs.
//   - time.Time values represent time instants.
//   - One time instant can be represented by many time.Time values.
//   - Comparing time should be done using the dedicated methods like `Equal`, `After`, `Before`, `Compare` or `IsZero`.
//   - time.Time can optionally contain a monotonic clock reading when created via time.Now()
//   - The monotonic clock reading is useful for measuring elapsed time.
//   - The monotonoc clock reading is not exposed.
//   - The monotonic clock reading is seperate from the system clock and only relevant inside the "current system boot" context.
//
// [How to parse a time or date in Go]: https://www.willem.dev/articles/how-to-parse-time-date/
// [time.Time and time.Location explained]: https://www.willem.dev/articles/time-location-explained/
// [time.Now and the monotonic clock]: https://www.willem.dev/articles/time-now-monotonic-clock/
// [Comparing a time or date in Go]: https://www.willem.dev/articles/how-to-compare-time-date/
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	p, err := time.Parse(time.RFC3339, "2024-01-18T13:37:00Z") // time.RFC3339 is a reference layout.
	if err != nil {
		log.Fatalf("failed to parse time: %v", err)
	}

	now := time.Now() // returns a time.Time in Local time.
	fmt.Println("p:", p)
	fmt.Println("now:", now) // the m=+0.0xxx is the monotonic time

	fmt.Println("now location:", now.Location())
	fmt.Println("p location:", p.Location())

	if now.After(p) {
		fmt.Println("now is after p")
	}
}

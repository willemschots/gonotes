// # Using context.Context
//
// This command explores context functionality. This is a small demo related to my articles on context.Context:
//
//   - [Context cancellation explained].
//   - [Choosing between context.TODO and context.Background].
//   - [How to add values to context].
//   - [What values to store in contexts].
//
// # Notes
//
//   - Customary to pass contexts as the first parameter to a function.
//   - Contexts form a tree.
//   - Contexts can be cancelled by timeout or by explicit cancellation.
//   - The cancellation is propagated to all child contexts.
//   - context.TODO() and context.Background() both return an empty context.
//   - Contexts can be used to store values.
//   - Contexts work with the empty interface any or interface{}. Useful to provide type-safe access functions.
//   - Best to only store developer relevant metadata in contexts. Things like request IDs etc.
//
// [Context cancellation explained]: https://www.willem.dev/articles/context-cancellation-explained/
// [Choosing between context.TODO and context.Background]: https://www.willem.dev/articles/context-todo-or-background/
// [How to add values to context]: https://www.willem.dev/articles/how-to-add-values-to-context/
// [What values to store in contexts]: https://www.willem.dev/articles/what-values-store-in-context/
package main

import (
	"context"
	"fmt"
)

type ctxKey string

func withValue(ctx context.Context, s string) context.Context {
	return context.WithValue(ctx, ctxKey("key"), s)
}

func getValue(ctx context.Context) (string, bool) {
	s, ok := ctx.Value(ctxKey("key")).(string)
	return s, ok
}

func doSomething(ctx context.Context) {
	s, ok := getValue(ctx)
	if !ok {
		return
	}

	fmt.Println(s)
}

func main() {
	// Create a new base context.
	ctx := context.Background()

	// Add a value to the context.
	ctx = withValue(ctx, "hello world")

	// Do something with the context.
	doSomething(ctx) // also reads the value.
}

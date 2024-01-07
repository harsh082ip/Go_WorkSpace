package main

import (
	"context"
	"fmt"
)

type contextKey string

func main() {

	var authToken contextKey = "auth_token"

	ctx := context.WithValue(context.Background(), authToken, "XYZ_123")
	fmt.Println(ctx.Value(authToken))
	fmt.Println(authToken)
}

package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent ...
type MyEvent struct {
	Name string `json:"name"`
}

// HandleRequest ...
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Println("name: ", name)
	fmt.Printf("ctx: %+v\n", ctx)

	for {
		fmt.Println("Just wasting time...")
	}

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}

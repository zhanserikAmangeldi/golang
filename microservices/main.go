package main

import (
	"context"
	"fmt"

	"github.com/zhanserikAmangeldi/golang/microservices/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start server: %w", err)
	}
}

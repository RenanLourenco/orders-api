package main

import (
	"context"
	"fmt"

	"github.com/RenanLourenco/orders-api/application"
)

func main(){
	ctx := context.Background()
	app := application.New()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start the app: ", err)
	}
}


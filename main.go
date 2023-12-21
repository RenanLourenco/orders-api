package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/RenanLourenco/orders-api/application"
)

func main(){

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := application.New()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start the app: ", err)
	}
}


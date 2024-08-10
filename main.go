package main

import (
	"context"
	"fmt"
	"hello/app"
)

func main()  {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil{
		fmt.Println("Failed to start server")
	}
}

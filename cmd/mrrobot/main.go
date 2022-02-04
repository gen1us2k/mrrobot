package main

import (
	"greeter_bot/app"
	"greeter_bot/config"
	"log"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	handler := app.NewHandler(c)
	log.Fatal(handler.Start())
}

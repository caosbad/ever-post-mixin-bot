package main

import (
	"context"
	"flag"
	"log"

	"github.com/caosbad/ever-post-mixin-bot/durable"
	"github.com/caosbad/ever-post-mixin-bot/services"
)

func main() {
	service := flag.String("service", "http", "run a service")
	flag.Parse()
	db := durable.OpenDatabaseClient(context.Background())
	defer db.Close()

	switch *service {
	case "http":
		err := StartServer(db)
		if err != nil {
			log.Println(err)
		}
	case "blaze":
		hub := services.NewHub(db)
		err := hub.StartService(*service)
		if err != nil {
			log.Println(err)
		}
	}
}

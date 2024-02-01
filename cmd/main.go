package main

import (
	"os"

	"github.com/uiansol/product-follow-up/internal/adapters/api"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "migrate" {
			api.MigrateDB()
		}
	} else {
		api.MigrateDB()
		api.RunServer()
	}
}

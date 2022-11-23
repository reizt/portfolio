package main

import (
	"log"
	"os"

	"github.com/reizt/portfolio/src/drivers/db"
	"github.com/reizt/portfolio/src/drivers/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("2nd argument is required")
		return
	}

	switch os.Args[1] {
	case "start":
		if len(os.Args) < 3 {
			log.Fatal("3rd argument is required")
			return
		}
		switch os.Args[2] {
		case "dev":
			log.SetFlags(log.Lshortfile)
			server.LaunchDevServer()
			return
		}
	case "migrate":
		db.MigrateSync()
		return
	}
}

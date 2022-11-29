package main

import (
	"log"
	"os"
	"strconv"

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
			port := 3000
			if len(os.Args) >= 5 && os.Args[3] == "-p" {
				maybePort, err := strconv.Atoi(os.Args[4])
				if err == nil {
					port = maybePort
				}
			}
			log.SetFlags(log.Lshortfile)
			server.LaunchDevServer(port)
			return
		}
	case "migrate":
		db.MigrateSync()
		return
	default:
		log.Fatal("unknown arg")
	}
}

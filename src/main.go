package main

import (
	"log"

	"auditlog/server"
)

func main() {
	log.Println("Main: Server starting...")
	server.Start()
	log.Println("Main: Server shutdown")
}

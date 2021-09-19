package main

import (
	httpserver "DevOpsAlarm/httpServer"
	"log"
)

func main() {
	log.Println("Starting DevOps Alarm...")

	httpserver.StartHTTPServer()
}

package main

import (
	httpserver "DevOpsAlarm/httpServer"
	"log"
)

func main() {
	log.Println("Starting DevOps Alarm...")

	// utils.HandleError(rpio.Open())
	// pin := rpio.Pin(27) // GPIO 27, see https://pinout.xyz/pinout/pin11_gpio17
	// pin.Output()

	// for {
	// 	pin.Toggle()
	// 	time.Sleep(500 * time.Millisecond)
	// }

	httpserver.StartHTTPServer()
}

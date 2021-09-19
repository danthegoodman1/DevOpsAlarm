package httpserver

import (
	"DevOpsAlarm/utils"
	"fmt"
	"log"
	"time"

	"github.com/brian-armstrong/gpio"
	"github.com/labstack/echo/v4"
)

var (
	AlarmActive    bool = false // default false
	LastAlarmStart time.Time
	Ticker         *time.Ticker
	AbortChan      chan bool
)

func Post_Alarm(c echo.Context) error {
	log.Println("Activating alarm!")
	ActivateAlarm()
	return c.String(200, "Activated alarm!")
}

func ActivateAlarm() {
	LastAlarmStart = time.Now()
	AlarmActive = true
}

func DeactivateAlarm() {
	AlarmActive = false
}

// Setup the ticker to run on an interval
func SetupTicker() {
	Ticker = time.NewTicker(500 * time.Millisecond)
	AbortChan = make(chan bool)
	go AlarmLoop()
}

func AlarmLoop() {
	pinState := false
	pin := gpio.NewOutput(13, pinState) // GPIO 27, see https://pinout.xyz/pinout/pin11_gpio17
	for {
		select {
		case <-AbortChan: // Exit the alarm loop
			return
		case <-Ticker.C: // Check for whether the alarm is active
			utils.DebugLog(fmt.Sprintf("Alarm less then 1 min old: %t", time.Now().Before(LastAlarmStart.Add(time.Minute*1))))
			utils.DebugLog(fmt.Sprintf("Alarm active: %t", AlarmActive))
			if time.Now().Before(LastAlarmStart.Add(time.Minute*1)) && AlarmActive { // Only last for 1 minute, or while enabled
				utils.DebugLog("Enabling alarm")
				// Toggle physical alarm pin
				if pinState {
					pinState = false
					pin.Low()
				} else {
					pinState = true
					pin.High()
				}
			} else {
				utils.DebugLog("Disabling alarm")
				AlarmActive = false
				pin.Low() // Turn off physical alarm
			}
		}
	}
}

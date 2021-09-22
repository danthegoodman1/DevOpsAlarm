package httpserver

import (
	"DevOpsAlarm/utils"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stianeikeland/go-rpio/v4"
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

func Get_Ack(c echo.Context) error {
	if AlarmActive {
		log.Println("Acknowledged, deactivating alarm")
		DeactivateAlarm()
	}
	return c.String(200, "Acked")
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
	utils.HandleError(rpio.Open())
	pin := rpio.Pin(27) // GPIO 27, see https://pinout.xyz
	pin.Output()
	for {
		select {
		case <-AbortChan: // Exit the alarm loop
			rpio.Close()
			return
		case <-Ticker.C: // Check for whether the alarm is active
			utils.DebugLog(fmt.Sprintf("Alarm less then 1 min old: %t", time.Now().Before(LastAlarmStart.Add(time.Minute*1))))
			utils.DebugLog(fmt.Sprintf("Alarm active: %t", AlarmActive))
			if time.Now().Before(LastAlarmStart.Add(time.Minute*1)) && AlarmActive { // Only last for 1 minute, or while enabled
				// Toggle physical alarm pin
				utils.DebugLog("Enabling alarm")
				pin.Toggle()
			} else {
				utils.DebugLog("Disabling alarm")
				AlarmActive = false
				pin.Low() // Turn off physical alarm
			}
		}
	}
}

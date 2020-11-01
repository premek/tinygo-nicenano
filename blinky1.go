package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"machine"
	"time"
)

func main() {
    led := machine.P0_15
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 200)

		led.High()
		time.Sleep(time.Millisecond * 2000)
	}
}

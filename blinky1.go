package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"machine"
	"time"
)

func main() {
	led := machine.GPIO{machine.ADC0}

	led.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)

	}
}

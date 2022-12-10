package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var (
	// Use mcu pin 16, corresponds to physical pin 36
	unlockPin = rpio.Pin(16)
)

func Setup() error {
	// Open and map memory to access gpio, check for errors
	err := rpio.Open()
	if err != nil {
		return err
	}

	// Set pin to output mode
	unlockPin.Output()

	return nil
}

func Close() {
	// Close the relay when exiting
	unlockPin.Low()

	// Unmap gpio memory when done
	rpio.Close()
}

var mutexUnlockingCh = make(chan struct{}, 1)

func Unlock() {
	select {
	case mutexUnlockingCh <- struct{}{}:
		unlockingTheDoor()
		stopUnlockingTheDoorAfterSeconds(time.Second * 3)
		return
	default:
		// lock not acquired, someone else is unlocking atm
		return
	}
}

func unlockingTheDoor() {
	unlockPin.High()
}

func stopUnlockingTheDoorAfterSeconds(interval time.Duration) {
	go func() {
		tick := time.Tick(interval)
		select {
		case <-tick:
			unlockPin.Low()
			<-mutexUnlockingCh
		}
	}()
}

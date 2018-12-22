package main

import (
	"strconv"
	"sync"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	// How long to keep the on/off relay energized
	onSwitchDelay = 250
)

// TrainSwitch models the circuit used to control the turnout switch
type TrainSwitch struct {
	onOffSwitch *gpio.RelayDriver
	trackSwitch *gpio.RelayDriver
	state       bool
	lock        sync.Mutex
}

// NewTrainSwitch creates a new turnout switch
func NewTrainSwitch(r *raspi.Adaptor, onOffPin, directionPin int) *TrainSwitch {
	ts := &TrainSwitch{
		onOffSwitch: gpio.NewRelayDriver(r, strconv.Itoa(onOffPin)),
		trackSwitch: gpio.NewRelayDriver(r, strconv.Itoa(directionPin)),
	}

	ts.onOffSwitch.Off()
	ts.trackSwitch.Off()

	return ts
}

// Status returns true if the switch is on
func (ts *TrainSwitch) Status() bool {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	return ts.state
}

// Toggle toggles a switch between off and on
func (ts *TrainSwitch) Toggle() (bool, error) {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	if err := ts.executeDirectional(!ts.state); err != nil {
		return false, err
	}

	return ts.state, nil
}

// On turns a switch to the On position
func (ts *TrainSwitch) On() error {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	return ts.executeDirectional(true)
}

// Off turns the switch off
func (ts *TrainSwitch) Off() error {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	return ts.executeDirectional(false)
}

func (ts *TrainSwitch) executeDirectional(on bool) error {
	if ts.state == on {
		return nil
	}

	f := ts.trackSwitch.Off
	if on {
		f = ts.trackSwitch.On
	}

	if err := f(); err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 20)

	if err := ts.onOffSwitch.On(); err != nil {
		return err
	}

	time.Sleep(time.Millisecond * onSwitchDelay)

	if err := ts.onOffSwitch.Off(); err != nil {
		return err
	}

	ts.state = on

	return nil
}

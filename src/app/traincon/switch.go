package main

import (
	"sync"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	onSwitchDelay = 250
)

type TrainSwitch struct {
	onOffSwitch *gpio.RelayDriver
	trackSwitch *gpio.RelayDriver
	state       bool
	lock        sync.Mutex
}

func NewTrainSwitch(r *raspi.Adaptor, onOffPin, trackPin string) *TrainSwitch {
	return &TrainSwitch{
		onOffSwitch: gpio.NewRelayDriver(r, onOffPin),
		trackSwitch: gpio.NewRelayDriver(r, trackPin),
	}
}

func (ts *TrainSwitch) Status() bool {
	return ts.state
}

func (ts *TrainSwitch) Toggle() error {
	var err error
	if ts.state {
		err = ts.Off()
	} else {
		err = ts.On()
	}

	if err != nil {
		return err
	}

	return nil
}

func (ts *TrainSwitch) On() error {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	if ts.state {
		return nil
	}

	if err := ts.trackSwitch.On(); err != nil {
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

	ts.state = true

	return nil

}

func (ts *TrainSwitch) Off() error {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	if !ts.state {
		return nil
	}

	if err := ts.trackSwitch.Off(); err != nil {
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

	ts.state = false

	return nil

}

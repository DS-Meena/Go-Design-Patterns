package main

// Abstractioin
type RemoteControl struct {
	device Device
}

func (r *RemoteControl) TogglePower() {
	r.device.TurnOn()
	r.device.TurnOff()
}

// Refined Abstraction
type AdvancedRemoteControl struct {
	RemoteControl
}

func (a *AdvancedRemoteControl) SetChannel(channel int) {
	a.device.SetChannel(channel)
}

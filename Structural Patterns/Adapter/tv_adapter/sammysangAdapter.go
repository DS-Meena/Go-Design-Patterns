package main

type sammysangAdapter struct {
	// TODO: add a field for the SammysangTV reference
	sstv *SammysangTV
}

func (ss *sammysangAdapter) turnOn() {
	// TODO
	ss.sstv.setOnState(true)
}

func (ss *sammysangAdapter) turnOff() {
	// TODO
	ss.sstv.setOnState(false)
}

func (ss *sammysangAdapter) volumeUp() int {
	// TODO
	vol := ss.sstv.getVolume() + 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *sammysangAdapter) volumeDown() int {
	// TODO
	vol := ss.sstv.getVolume() - 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *sammysangAdapter) channelUp() int {
	ch := ss.sstv.getChannel() + 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *sammysangAdapter) channelDown() int {
	ch := ss.sstv.getChannel() - 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *sammysangAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
}
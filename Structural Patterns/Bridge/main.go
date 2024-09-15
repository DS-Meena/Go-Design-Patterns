package main

import "fmt"

func main() {
	tv := &TV{}
	radio := &Radio{}

	tvRemote := &RemoteControl{device: tv}
	radioRemote := &AdvancedRemoteControl{RemoteControl{device: radio}}

	fmt.Println("Testing TV remote: ")
	tvRemote.TogglePower()

	fmt.Println("\nTesting Radio remote:")
	radioRemote.TogglePower()
	radioRemote.SetChannel(100)
}

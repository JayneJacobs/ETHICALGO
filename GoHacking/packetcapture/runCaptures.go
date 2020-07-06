package main

import (
	"JayneJacobs/packetcapture/capture"
	"JayneJacobs/packetcapture/finddev"
)

func main() {

	finddev.FindDevices()
	capture.PacketNeat()
}

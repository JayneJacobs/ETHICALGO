package finddev

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

// FindDevices lists the interfaces for the network and other devices
func FindDevices() {

	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println(err)
	}

	for index, dev := range devices {
		fmt.Println(index, " ", dev.Name)
		for _, add := range dev.Addresses {
			fmt.Println("\tIP", add.IP)
			fmt.Println("\tNM", add.Netmask)

		}
	}

}

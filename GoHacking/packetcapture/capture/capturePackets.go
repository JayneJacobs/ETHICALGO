package capture

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface    = "eth0"
	devFound = false
	snaplen  = int32(1600)
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	promisc  = false
)

// PacketsRaw captures packets to a shell
func PacketsRaw() ([]pcap.Interface, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic(err)
	}

	for _, dev := range devices {
		if dev.Name == iface {
			devFound = true
		}
		if !devFound {
			log.Panicf("device not found %s", iface)
		}
		handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Panicln(err)
		}
		defer handle.Close()
		err = handle.SetBPFFilter(filter)
		if err != nil {
			log.Panicln(err)
		}

		// src := gopacket.NewPacketSource(handle, handle.LinkType())

		// for pkt := range src.Packets() {
		// 	fmt.Println(pkt)
		// }
	}
	return devices, err
}

// PacketNeat extracts username password Https is encrypted so this will only work for unencrypted
func PacketNeat() {
	devices, err := PacketsRaw()
	if err != nil {
		log.Panic(err)
	}
	for _, dev := range devices {
		if dev.Name == iface {
			devFound = true
		}
		if !devFound {
			log.Panicf("device not found %s", iface)
		}
		handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Panicln(err)
		}
		defer handle.Close()
		err = handle.SetBPFFilter(filter)
		if err != nil {
			log.Panicln(err)
		}

		src := gopacket.NewPacketSource(handle, handle.LinkType())

		for pkt := range src.Packets() {
			applayer := pkt.ApplicationLayer()
			if applayer == nil {
				continue
			}
			payload := applayer.Payload()
			searchArr := []string{"jayne", "Jayne", "Diptera"}
			for _, s := range searchArr {
				index := strings.Index(string(payload), s)
				if index != -1 {
					fmt.Println(string(payload[index : index+100]))
				}
			}
		}

	}
}

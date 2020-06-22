package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	targetIP := "192.168.221.1/24"

	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targetIP),
		nmap.WithPorts("80, 443"),
		nmap.WithContext(ctx),
	)

	if err != nil {
		log.Fatal(err)
	}

	results, warning, err := scanner.Run()

	if err != nil {
		log.Fatal(err)
	}

	if warning != nil {
		log.Fatalf("warning %s \n", warning)
	}

	for _, host := range results.Hosts {
		if len(host.Ports) == 0 {
			if len(host.Ports) == 0 || len(host.Addresses) == 0 {
				continue
			}

			fmt.Printf("IP: %q\n", host.Addresses)

			if len(host.Addresses) > 1 {
				fmt.Printf("IP: %q", host.Addresses[0])
			}

			for _, port := range host.Ports {
				fmt.Printf("\t Port %d %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			}

		}
	}

}

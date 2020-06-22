package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string, argsarr []string) (err error) {
	args := argsarr
	cmdobj := exec.Command(command, args...)

	cmdobj.Stdout = os.Stdout
	cmdobj.Stderr = os.Stderr
	cmdobj.Stdin = os.Stdin

	err = cmdobj.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	command := "sudo"
	// 00:0c:29:2a:ad:39
	//  ifconfig eth0 down
	//  ifconfig eth hw ether 00:11:22:33:44:55
	//  ifconfig eth0 up
	// oldMAC := "00:0c:29:2a:ad:39"
	iface := flag.String("iface", "", "Interface for which you want to change MAC.")
	newMAC := flag.String("newMAC", "", "provide teh new mac address")
	flag.Parse()
	executeCommand(command, []string{"ifconfig", *iface, "down"})
	executeCommand(command, []string{"ifconfig", *iface, "hw", "ether", *newMAC})
	executeCommand(command, []string{"ifconfig", *iface, "up"})
	executeCommand(command, []string{"ifconfig"})
	// executeCommand(command, []string{"ifconfig", *iface, "down"})
	// executeCommand(command, []string{"ifconfig", *iface, "hw", "ether", oldMAC})
	// executeCommand(command, []string{"ifconfig", *iface, "up"})
	// executeCommand(command, []string{"ifconfig"})
}

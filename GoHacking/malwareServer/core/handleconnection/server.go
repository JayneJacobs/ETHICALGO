package handleconnection

import (
	"net"
)

// ConnecttoVictim takes and IP address and Port and returns a connection and error.
func ConnecttoVictim(IP string, Port string) (connection net.Conn, err error) {
	LocalAddress := IP + ":" + Port

	listener, err := net.Listen("tcp", LocalAddress)

	if err != nil {
		return connection, err
	}

	connection, err = listener.Accept()

	if err != nil {
		return connection, err
	}

	return
}

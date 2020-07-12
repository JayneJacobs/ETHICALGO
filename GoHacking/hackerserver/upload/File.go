package upload

import (
	"bufio"
	"encoding/gob"
	"errors"
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

// ReadFile (filename string) ([]byte, error)
func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[+] Unable to oopen file ")
	}
	defer file.Close()
	stats, err := file.Stat()
	FileSize := stats.Size()
	fmt.Printf("[+] The file contains, %v bytes \n  ", FileSize)

	bytes := make([]byte, FileSize)
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// FiletoRemote (connection net.Conn) (err error)
func FiletoRemote(connection net.Conn) (err error) {
	filename := "file.png"
	fileExists := CheckExistence(filename)
	if !fileExists {
		err := errors.New("File not found")
		return err
	}
	content, err := ReadFile(filename)
	filesize := len(content)
	fmt.Println("File does exist?", fileExists)
	fs := &SendFile{
		Filename:    filename,
		FileSize:    filesize,
		FileContent: nil,
	}

	encoder := gob.NewEncoder(connection)
	err = encoder.Encode(fs)
	if err != nil {
		return
	}

	reader := bufio.NewReader(connection)
	reader.ReadString('\n')
	status, err := reader.ReadString('\n')
	fmt.Println(status)
	return

}

// ln -s /usr/local/go/bin/go  /usr/sbin/go

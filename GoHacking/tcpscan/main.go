package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	IP := "scanme.nmap.org"
	// Port := "80"

	// address := IP + ":" + Port
	// connection, err := net.Dial("tcp", address)
	// if err == nil {
	// 	fmt.Println("+ Connection established", connection.RemoteAddr().String())
	// }
	var wg sync.WaitGroup
	for i := 20; i < 100; i++ {
		wg.Add(1)
		go scanMultiple(i, IP, &wg)
		wg.Wait()
	}
}

func scanMultiple(port int, IP string, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf(IP+":%d", port)
	fmt.Println(address)
	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
		// connection.Close
	}

	fmt.Printf("+ Connection established for %d, %v\n", port, connection.RemoteAddr().String())
	// connection.Close()

	connection.Close()

}

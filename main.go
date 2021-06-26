package main

import (
	"fmt"
	"log"
	"net"
)

// Get preferred outbound ip of this machine
func main() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Printf("%s\n",localAddr.IP)
}

package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil { //if connection has no errors, it worked
		fmt.Println("Connection successful")
	}
}

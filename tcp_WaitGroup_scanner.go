package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1) // create a thread and wait for finish
		go func(j int) {
			defer wg.Done() // wait for task completion before kill
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return // closed
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() // wait for all threads to finish
}

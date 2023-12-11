package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// https://www.youtube.com/watch?v=c6DH-1nffTI

// Listen for the time with telnet localhost 9000

func main() {

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func() {
			for {
				_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
				if err != nil {
					return
				}

				time.Sleep(time.Second)
			}
		}()
	}
}

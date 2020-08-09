package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	cancel := make(chan int)
	go func() {
		var command string
		for {
			fmt.Fscan(os.Stdin, &command) // read a command
			if strings.EqualFold(command, "exit") {
				cancel <- 1
			}
		}
	}()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				return
			}
			go handleConn(conn)

		}

	}()

	for {
		select {
		case <-cancel:
			fmt.Println("Server stoped!")
			return
		}
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)
	str := ""
	go func() {
		defer f.Close()
		defer close(out)
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]

			// If new line present in 8 byte window
			if i := bytes.IndexByte(data, '\n'); i != -1 {

				// Append data until the new line
				str += string(data[:i])

				// Move window to data after new line
				data = data[i+1:]

				out <- str

				// Reset str for next line
				str = ""
			}

			// Append data to str - either the entire window or the part after new line
			str += string(data)
		}

		if len(str) != 0 {
			out <- str
		}
	}()

	return out
}

func main() {

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("read: %s\n", line)
		}
	}
}

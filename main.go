package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	str := ""
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

			// Display str so far
			fmt.Printf("read: %s\n", str)

			// Reset str for next line
			str = ""
		}

		// Append data to str - either the entire window or the part after new line
		str += string(data)
	}
}

package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"fmt"
)

var data = make(map[string]string)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		// skip blank lines
		if len(fs) < 2 {
			continue
		}

		switch fs[0] {
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)
		case "SET":
			if len(fs) != 3 {
				io.WriteString(conn, "EXPECTED VALUE\n")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value
		case "DEL":
			key := fs[1]
			delete(data, key)
		default:
			io.WriteString(conn, "INVALID COMMAND "+fs[0]+"\n")
		}
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
		// ONLY HANDLES ONE CONNECTION AT A TIME
		// Could we make it concurrent? How?
		// What are the considerations?
	}
}

package main

import (
	"fmt"
	"net"
	"os"
)

var log = fmt.Println
var logf = fmt.Printf

func main() {
	conn, err := net.Dial("tcp", "localhost:8787")
	if err != nil {
		log("net dial failed")
		os.Exit(1)
	}
	message := "hihi"
	n, err := conn.Write([]byte(message))
	if err != nil {
		log("conn write failed")
		panic(err)
	}
	log("n: ", n)

	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		log("conn read failed")
		panic(err)
	}
	log("n: ", n)
	log("from server: ", string(buf))
}

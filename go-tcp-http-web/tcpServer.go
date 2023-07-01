package main

import "net"

func RunTcpServer(listener net.Listener) {

	for {
		conn, err := listener.Accept()
		if err != nil {
			log("conn accept failed ")
		}

		func(conn net.Conn) {
			defer conn.Close()

			// read from socket
			buf := make([]byte, 1024)
			_, err := conn.Read(buf)
			if err != nil {
				log("conn read failed")
			}
			log("from client: ", string(buf))

			// write to socket
			_, err = conn.Write(buf)
			if err != nil {
				log("conn write failed")
				panic(err)
			}
		}(conn)
	}
}

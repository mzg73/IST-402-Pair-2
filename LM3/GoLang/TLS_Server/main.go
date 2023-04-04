package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	// If the certificate is not nil (GoLang's version of null), print out a log error
	if err != nil {
		log.Println(err)
		return
	}

	// Listens for TLS connections on port 443 (protocol is tcp)
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":443", config)
	// If tcp listener is not nil (GoLang's version of null), print out a log error
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		// Starts actively listening for incoming connections (execute the "client.go" file to continue)
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

// Method that starts on its own thread (can handle connections concurrently
func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

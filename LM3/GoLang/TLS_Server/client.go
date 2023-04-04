package main

import (
	"crypto/tls"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		// Bypasses verification of the server's certificate.
		InsecureSkipVerify: true,
	}

	//  Dials the TCP server at "127.0.0.1:443" (Interestingly enough, 127.0.0.1 is a loopback address)
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Sends "hello" to the server
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	// Reads the sent message of "world" in a format of bytes
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	// Ends up printing out "world"
	println(string(buf[:n]))
}

package server

import (
	"log"
	"net"
	"testing"
	"time"
)

func MockClient() {

	log.Printf("[client] mock client starting ...\n")

	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:60000")

	if err != nil {
		log.Printf("[client] mock client start err: %s\n", err.Error())
		return
	}

	for {
		_, err := conn.Write([]byte("[client] Hello World"))

		if err != nil {
			log.Printf("[client] mock client write error: %s\n", err.Error())
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)

		if err != nil {
			log.Printf("[client] mock client read error: %s\n", err.Error())
			return
		}

		log.Printf("[client] server call back : %s, cnt = %d\n", string(buf), cnt)

		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("tcp_server", "tcp4", "0.0.0.0", 60000)

	go MockClient()

	s.Run()
}

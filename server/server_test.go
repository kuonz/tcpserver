package server

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func MockClient() {

	log.Println("[client] mock client starting ...")

	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:60000")

	if err != nil {
		fmt.Println("[client] mock client start err: ", err.Error())
		return
	}

	for {
		_, err := conn.Write([]byte("[client] Hello World"))

		if err != nil {
			fmt.Println("[client] mock client write error: ", err.Error())
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Println("[client] mock client read error: ", err.Error())
			return
		}

		fmt.Printf("[client] server call back : %s, cnt = %d", string(buf), cnt)

		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("tcp_server", "tcp4", "0.0.0.0", 60000)

	go MockClient()

	s.Run()
}

package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Server ...
type Server struct {
	Name      string // 服务器的名称
	IPVersion string // 服务器支持的IP版本号
	IPAddress string // 服务器监听的IP地址
	Port      uint   // 服务器监听的端口号
}

// NewServer ...
func NewServer(name, ipversion, ipaddress string, port uint) *Server {
	return &Server{
		Name:      name,
		IPVersion: ipversion,
		IPAddress: ipaddress,
		Port:      port,
	}
}

// Run ...
func (s *Server) Run() {

	var wg sync.WaitGroup

	wg.Add(1)

	s.Serve()

	log.Printf("[server] server has started ...\n")

	wg.Wait()
}

// Serve ...
func (s *Server) Serve() {

	// 创建套接字
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddress, s.Port))

	if err != nil {
		log.Printf("[server] server serve failed, err: %s\n", err.Error())
		return
	}

	// 监听端口，获取 listener
	listener, err := net.ListenTCP(s.IPVersion, addr)

	if err != nil {
		log.Printf("[server] liesten err: %s\n", err.Error())
		return
	}

	// 循环监听，获取连接，处理连接
	go func() {
		for {
			conn, err := listener.AcceptTCP()

			if err != nil {
				log.Printf("[server] accept tcp failed, err: %s\n", err.Error())
				continue
			}

			go temp(conn)
		}
	}()
}

// Stop ...
func (s *Server) Stop() {
	log.Printf("[server] server stop, release resources ...\n")
}

func temp(conn *net.TCPConn) {
	for {
		buf := make([]byte, 512)

		cnt, err := conn.Read(buf)

		if err != nil {
			log.Printf("[server] process tcp conn failed, err: %s\n", err.Error())
			return
		}

		log.Printf("[server] server received conn, data: %s\n", string(buf[:cnt]))

		_, err = conn.Write(buf[:cnt])

		if err != nil {
			log.Printf("server response write failed, err: %s\n", err.Error())
		}
	}
}

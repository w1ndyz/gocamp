package main

import (
	"bufio"
	"fmt"
	"net"
)

func InitTCP(addr string) error {
	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", addr); err != nil {
		fmt.Printf("net.Listen(tcp, %s) err: %v", addr, err)
		return err
	}

	acceptTCP(listener)
	return nil
}

func acceptTCP(listen net.Listener) {
	var (
		conn net.Conn
		err  error
	)
	for {
		if conn, err = listen.Accept(); err != nil {
			fmt.Printf("listener.Accept err: %v", err)
			return
		}
		go serveTCP(conn)
	}
}

func serveTCP(conn net.Conn) {
	defer conn.Close()
	p := &Proto{}
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)

	for {
		err := p.ReadTCP(rd)
		if err != nil {
			fmt.Printf("read error: %v", err)
			return
		}
		fmt.Printf("seq:[%d], op:[%d], body:[%s]\n", p.Seq, p.Op, string(p.Body))
		err = p.WriteTCP(wr)
		if err != nil {
			fmt.Printf("write error [%v]", err)
			return
		}
	}
}

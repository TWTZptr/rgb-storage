package api

import "net"

type Response struct {
	Val string
	Err string
}

type ResponseWriter interface {
	Write(net.Conn, Response) error
}

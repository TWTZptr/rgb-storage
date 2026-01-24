package api

import "net"

type Response struct {
	Err string
}

type ResponseWriter interface {
	Write(net.Conn, Response) error
}

package main

import (
	"net"
	"strings"
)

func initRequest(connection net.Conn) *request {
	req := request{
		connection: connection,
	}
	buf := make([]byte, 1024)
	length, _ := connection.Read(buf)
	data := string(buf[:length])
	lines := strings.Split(data, "\r\n")
	line1 := strings.Split(lines[0], " ")
	req.method = line1[0]
	req.path = line1[1]
	req.version = line1[2]
	req.headers = make(map[string]string)
	for _, v := range lines[1:] {
		if v == "\r\n" {
			break
		}
		key, value, ok := strings.Cut(v, ":")
		if ok {
			req.headers[key] = value
		}
	}
	return &req
}

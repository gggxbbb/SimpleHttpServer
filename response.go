package main

import (
	"net"
	"strconv"
	"strings"
	"time"
)

func initResponse(connection net.Conn) *response {
	rep := response{
		connection: connection,
		version:    "HTTP/1.1",
		code:       200,
		status:     "OK",
		headers: map[string]string{
			"Content-Type": "text/html, charset=utf-8",
			"Date":         time.Now().UTC().Format("Mon, 04 Jan 2006 15:04:05 GMT"),
			"Server":       "SimpleHttpServer",
		},
		body: []byte(""),
	}
	return &rep
}

func (rep response) Send() {
	data := []byte(
		rep.version + " " + strconv.Itoa(rep.code) + " " + rep.status + "\n\r")
	for k, v := range rep.headers {
		data = append(data, []byte(strings.ToUpper(k)+": "+v+"\r\n")...)
	}
	data = append(data, []byte("Content-Length: "+strconv.Itoa(len(rep.body)))...)
	data = append(data, []byte("\r\n\r\n")...)
	data = append(data, rep.body...)
	_, err := rep.connection.Write(data)
	if err != nil {
		return
	}
}

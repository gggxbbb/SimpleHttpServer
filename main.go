package main

import (
	"net"
)
import "fmt"

func main() {
	listener, err := net.Listen("tcp", ":10080")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		req := initRequest(connection)
		fmt.Printf("->: %s %s %s \n", req.version, req.method, req.path)
		rep := initResponse(connection)
		if req.method != "GET" {
			rep.code = 500
			rep.status = "ERROR"
			rep.WriteBody([]byte("Only Support GET Method"))
		} else {
			rep.WriteBody([]byte("Hello Word"))
		}
		fmt.Printf("<-: %s %d %s \n", rep.version, rep.code, rep.status)
		rep.Send()
		if err != nil {
			return
		}
		err = connection.Close()
		if err != nil {
			return
		}
	}

}

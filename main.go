package main

import (
	"net"
)
import "fmt"

func main() {
	listener, err := net.Listen("tcp", ":10080")

	routes = make(map[string]onReq)

	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	initRoutes()

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
			rep.Send()
		} else {
			requestHandler(req, rep)
		}
		fmt.Printf("<-: %s %d %s \n", rep.version, rep.code, rep.status)
		if err != nil {
			return
		}
		err = connection.Close()
		if err != nil {
			return
		}
	}

}

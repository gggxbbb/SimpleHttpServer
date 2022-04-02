package main

import "net"

type request struct {
	connection net.Conn
	method     string
	path       string
	version    string
	headers    map[string]string
	body       []byte
}

type response struct {
	connection net.Conn
	version    string
	code       int
	status     string
	headers    map[string]string
	body       []byte
}

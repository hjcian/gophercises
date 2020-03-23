package main

import "fmt"

type client struct{}

func newConnect(host string, port int, table string) (*client, error) {
	if host == nil {
		host = "127.0.0.1"
	}
	return &client{}, nil
}

func main() {
	c, err := newConnect("127.0.0.1", 5432, "tx")
	// c, err := newConnect("127.0.0.1", 5432)
	// c, err := newConnect("127.0.0.1")
	fmt.Println(c, err)
}

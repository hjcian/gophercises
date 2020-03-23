package main

import (
	"errors"
	"fmt"
)

type dbConfig struct {
	Host  string
	Port  int
	Table string
}

type client struct {
	*dbConfig
}

func newConnect(conf *dbConfig) (*client, error) {
	// do some necessary process
	return &client{dbConfig: conf}, nil
}

type dbConfigOption func(c *dbConfig) error

func setHost(host string) dbConfigOption {
	return func(c *dbConfig) error {
		if c == nil {
			return errors.New("given config is nil")
		}
		c.Host = host
		return nil
	}
}

func setPort(port int) dbConfigOption {
	return func(c *dbConfig) error {
		if c == nil {
			return errors.New("given config is nil")
		}
		c.Port = port
		return nil
	}
}

func setTable(table string) dbConfigOption {
	return func(c *dbConfig) error {
		if c == nil {
			return errors.New("given config is nil")
		}
		c.Table = table
		return nil
	}
}

func wrapConnect(options ...dbConfigOption) (*client, error) {
	// default configuration
	conf := &dbConfig{
		Host:  "127.0.0.1",
		Port:  5432,
		Table: "tx",
	}
	for _, f := range options {
		err := f(conf)
		if err != nil {
			// do some error handling
			fmt.Println(err)
		}
	}
	return newConnect(conf)
}

func main() {
	db, err := wrapConnect(setHost("192.168.1.130"), setPort(1234))
	// db, err := wrapConnect(setHost("192.168.1.130"))
	// db, err := wrapConnect()
	fmt.Println(db.Host)
	fmt.Println(db.Port)
	fmt.Println(db.Table)
	fmt.Println(err)
}

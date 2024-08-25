package client

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Connection struct {
	Conn net.Conn
	Id   int
	Ip   string
	Port int
}

func (c *Connection) Connect() error {
	err := c.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		log.Fatalf("SetReadDeadline failed: %s", err)
		// do something else, for example create new conn
		return fmt.Errorf("SetReadDeadline failed: %s", err)
	}

	err = c.Conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatalf("SetWriteDeadline failed: %s", err)
		// do something else, for example create new conn
		return fmt.Errorf("SetWriteDeadline failed: %s", err)
	}

	return nil
}

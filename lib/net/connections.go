package net

import (
	"fmt"
)

type Connections struct {
	Connections []Connection
}

func (c *Connections) AddConnection(conn Connection) {
	fmt.Printf("Adding connection: %v\n", conn)
	c.Connections = append(c.Connections, conn)
}

func (c *Connections) RemoveConnection(id int) {
	for i, conn := range c.Connections {
		if conn.Id == id {
			fmt.Printf("Removing connection: %v\n", conn)
			c.Connections = append(c.Connections[:i], c.Connections[i+1:]...)
		}
	}
}

func (c *Connections) GetConnection(id int) (connection Connection, err error) {
	for _, conn := range c.Connections {
		if conn.Id == id {
			fmt.Printf("Getting connection: %v\n", conn)
			return conn, nil
		}
	}
	return Connection{}, fmt.Errorf("Connection not found")
}

func (c *Connections) HasConnection(id int) bool {
	for _, conn := range c.Connections {
		if conn.Id == id {
			return true
		}
	}
	return false
}

func (c *Connections) GetConnections() []Connection {
	return c.Connections
}

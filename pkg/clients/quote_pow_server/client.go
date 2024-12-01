package quote_pow_server

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	address string
	conn    net.Conn
}

func NewClient(port string) (*Client, error) {
	address := fmt.Sprintf("localhost:%s", port) // assuming localhost for simplicity
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	return &Client{
		address: port,
		conn:    conn,
	}, nil
}

func (c *Client) Send(message string) error {
	_, err := fmt.Fprintf(c.conn, "%s\n", message)
	return err
}

func (c *Client) Receive() (string, error) {
	reader := bufio.NewReader(c.conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read from server: %w", err)
	}
	return message, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

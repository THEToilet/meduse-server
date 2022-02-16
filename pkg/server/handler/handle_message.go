package handler

import (
	"net"
	"time"
)

func (c *Connection) handleMessage(rawMessage []byte, pongTimer *time.Timer) {
	functionType := rawMessage[0]
	print(functionType)

}

func (c *Connection) sendMessage(conn *net.Conn, message []byte) error {
	c.logger.Debug().Caller().Msg(string(message))
	n, err := c.conn.Write(message)
	c.logger.Info().Str("", "d").Msg(string(n))
	if err != nil {
		c.logger.Err(err).Msg(err.Error())
	}
	return nil
}

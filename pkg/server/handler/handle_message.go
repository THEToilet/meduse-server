package handler

import (
	"net"
	"time"
)

func (c *Connection) handleMessage(rawMessage []byte, pongTimer *time.Timer) {
	functionType := rawMessage[0]
	print(functionType)
	switch functionType {
	case 1:
		// NOTE: ユーザ登録リクエスト | 0x01
	case 2:
		// NOTE: 部屋一覧取得リクエスト（定期的） | 0x02
	case 3:
		// NOTE: 部屋参加リクエスト | 0x03
	case 4:
		// NOTE: パケット中継 | 0x04 | 2byte
	case 5:
		// NOTE: 部屋離脱 | 0x05
	case 6:
		// NOTE: アプリ離脱 | 0x06
	case 7:
		// NOTE: pong | 0x07
		stopTimer(pongTimer)
		pongTimer.Reset(time.Second * 10)
		c.logger.Info().Msg("pong")
		c.logger.Info().Interface("PONG", functionType).Interface("userID", c.user.UserID).Msg("RECEIVE-PONG-MESSAGE")
	default:
		c.logger.Info().Msg("Invalid RequestType")
	}

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

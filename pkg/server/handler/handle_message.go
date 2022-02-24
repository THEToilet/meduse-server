package handler

import (
	"encoding/json"
	"net"
	"strconv"
	"time"
)

func (c *Connection) handleMessage(rawMessage []byte, pongTimer *time.Timer) {
	functionType := rawMessage[0]
	print(functionType)
	switch functionType {
	case 1:
		// NOTE: pong | 0x01
		stopTimer(pongTimer)
		pongTimer.Reset(time.Second * 10)
		message := []byte("1")
		err := c.sendMessage(c.conn, message)
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}
		c.logger.Info().Msg("pong")
		c.logger.Info().Interface("PONG", functionType).Interface("userID", c.user.UserID).Msg("RECEIVE-PONG-MESSAGE")
	case 2:
		// NOTE: ユーザ登録リクエスト | 0x02
		userID, err := c.roomUseCase.RegisterUser()
		if err != nil {
			message := []byte("3" + userID)
			err := c.sendMessage(c.conn, message)
			if err != nil {
				c.logger.Fatal().Err(err).Msg("d")
			}
		}
		message := []byte("3" + userID)
		err = c.sendMessage(c.conn, message)
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}

		c.user.UserID = userID

		err = c.connections.Save(userID, &c.conn)
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}
	case 3:
		// NOTE: ユーザ一覧取得 | 0x03
		userList := c.roomUseCase.GetAllUser()
		userListByte, err := json.Marshal(userList)
		if err != nil {
			c.logger.Fatal().Err(err)
		}
		message := []byte("4")
		for _, v := range userListByte {
			message = append(message, v)
		}
		err = c.sendMessage(c.conn, message)
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}
	case 4:
		// NOTE: パケット中継 | 0x04 | 2byte
		host, err := c.roomUseCase.GetHostUser()
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}
		hostConn, err := c.connections.Find(host.UserID)
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}
		message := []byte("5" + string(rawMessage[1:]))
		err = c.sendMessage(*hostConn, message)
		if err != nil {
			c.logger.Fatal().Err(err).Msg("d")
		}
	case 5:
		// NOTE: アプリ離脱 | 0x05
		err := c.roomUseCase.DeleteUser(c.user.UserID)
		if err != nil {
			message := []byte("2" + err.Error())
			err := c.sendMessage(c.conn, message)
			if err != nil {
				c.logger.Fatal().Err(err).Msg("d")
			}
			c.logger.Fatal().Err(err).Msg("d")
		}
	case 6:
		// NOTE: 観客席に送る | 0x06
		c.logger.Info().Interface("uuid", rawMessage[1:17])
		err := c.roomUseCase.ReturnToSpectators(string(rawMessage[1:17]))
		if err != nil {
			message := []byte("2" + err.Error())
			err := c.sendMessage(c.conn, message)
			if err != nil {
				c.logger.Fatal().Err(err).Msg("d")
			}
			c.logger.Fatal().Err(err).Msg("d")
		}
	case 7:
		// NOTE: ユーザBAN | 0x07
		c.logger.Info().Interface("uuid", rawMessage[1:17])
		err := c.roomUseCase.DeleteUser(string(rawMessage[1:17]))
		if err != nil {
			message := []byte("2" + err.Error())
			err := c.sendMessage(c.conn, message)
			if err != nil {
				c.logger.Fatal().Err(err).Msg("d")
			}
			c.logger.Fatal().Err(err).Msg("d")
		}
	case 8:
		// NOTE: ユーザをコントローラに登録 | 0x08
		c.logger.Info().Interface("uuid", rawMessage[1:17])
		conNum, _ := strconv.Atoi(string(rawMessage[18]))
		err := c.roomUseCase.SpectatorMoveToController(string(rawMessage[1:17]), conNum)
		if err != nil {
			message := []byte("2" + err.Error())
			err := c.sendMessage(c.conn, message)
			if err != nil {
				c.logger.Fatal().Err(err).Msg("d")
			}
			c.logger.Fatal().Err(err).Msg("d")
		}
	default:
		c.logger.Info().Msg("Invalid RequestType")
	}
}

func (c *Connection) sendMessage(conn net.Conn, message []byte) error {
	c.logger.Debug().Caller().Msg(string(message))
	n, err := conn.Write(message)
	c.logger.Info().Str("", "d").Msg(string(n))
	if err != nil {
		c.logger.Err(err).Msg(err.Error())
	}
	return nil
}

package handler

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"github.com/rs/zerolog"
	"meduse-server/pkg/domain/application"
	"meduse-server/pkg/domain/model"
	"net"
	"time"
)

type Connection struct {
	conn           net.Conn
	receiveMessage chan []byte
	sendingMessage chan []byte
	connections    Connections
	roomUseCase    *application.RoomUseCase
	user           model.User
	logger         *zerolog.Logger
}

func NewConnection(conn net.Conn, receiveMessage chan []byte, sendingMessage chan []byte, connections Connections, roomUseCase *application.RoomUseCase, logger *zerolog.Logger) *Connection {
	return &Connection{
		conn:           conn,
		receiveMessage: receiveMessage,
		sendingMessage: sendingMessage,
		connections:    connections,
		roomUseCase:    roomUseCase,
		logger:         logger,
	}
}

func stopTimer(timer *time.Timer) {
	if !timer.Stop() {
		<-timer.C
	}
}

func (c *Connection) Selector(ctx context.Context, cancel context.CancelFunc) {
	pingTimer := time.NewTicker(10 * time.Second)
	pongTimer := time.NewTimer(10 * time.Second)
	stopTimer(pongTimer)
	defer func() {
		// TODO　ここ呼ばれるか確認
		stopTimer(pongTimer)
		pingTimer.Stop()
		c.logger.Debug().Caller().Msg("selector is close")
		if err := c.roomUseCase.DeleteUser(c.user.UserID); err != nil {
			c.logger.Debug().Msg("Delete error")
		}
		if err := c.connections.Delete(c.user.UserID); err != nil {
			c.logger.Debug().Msg("Connections Delete error")
		}
	}()

	// TODO 疎通確認のping pong
	// Labeled Break
L:
	for {
		select {
		case <-pingTimer.C:
			message := []byte{1}
			c.logger.Info().Interface("sendData", message).Interface("sendDataHex", hex.EncodeToString(message)).Interface("sendBinaryData", binary.Size(message)).Interface("userID", c.user.UserID).Msg("SEND-PONG-LOG")
			if err := c.sendMessage(c.conn, message); err != nil {
				c.logger.Error().Msg(err.Error())
				break L
			}
			pongTimer.Reset(10 * time.Second)
		case <-pongTimer.C:
			c.logger.Info().Msg("pong is failed")
			break L
		case msg, ok := <-c.receiveMessage:
			if !ok {
				c.logger.Fatal().Msg("d")
				break L
			}
			c.handleMessage(msg, pongTimer)
		case msg, ok := <-c.sendingMessage:
			c.logger.Info().Caller().Msg(string(msg))
			if !ok {
				c.logger.Fatal().Msg("")
				break L
			}
			c.logger.Info().Interface("sendData", msg).Interface("sendBinaryData", binary.Size(msg)).Interface("userID", c.user.UserID).Msg("SEND-MESSAGE-LOG")
			// クライアントへメッセージ送信
			if err := c.sendMessage(c.conn, msg); err != nil {
				break L
			}
		}
	}
	// TODO: サーバが死んだことによる
	// TODO: エラーメッセージをクライアント側へ送る
	if err := c.sendMessage(c.conn, []byte{}); err != nil {
		c.logger.Debug().Msg("")
		return
	}
	cancel()
}

func (c *Connection) Receiver(ctx context.Context) {
	defer func() {
		c.logger.Debug().Msg("Socket is closed")
		c.conn.Close()
	}()
L:
	for {
		data := make([]byte, 128)
		length, err := c.conn.Read(data)
		if err != nil {
			c.logger.Err(err).Msg(err.Error())
			break L
		}
		c.logger.Info().Interface("data", data).Msg("receive raw data")
		c.logger.Info().Interface("data", hex.EncodeToString(data)).Msg("receive hex data")

		if length == 0 {
			c.logger.Info().Str("", "d").Msg("connection is closed")
			break L
		} else {
			c.receiveMessage <- data
		}
	}
	<-ctx.Done()
}

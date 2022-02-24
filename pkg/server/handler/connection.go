package handler

import (
	"context"
	"encoding/binary"
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
	roomUseCase    *application.RoomUseCase
	user           model.User
	logger         *zerolog.Logger
}

func NewConnection(conn net.Conn, receiveMessage chan []byte, sendingMessage chan []byte, roomUseCase *application.RoomUseCase, logger *zerolog.Logger) *Connection {
	return &Connection{
		conn:           conn,
		receiveMessage: receiveMessage,
		sendingMessage: sendingMessage,
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
		if err := c.roomUseCase.Delete(c.user.UserID); err != nil {
			c.logger.Debug().Msg("Delete error")
		}
	}()

	// TODO 疎通確認のping pong
	// Labeled Break
L:
	for {
		select {
		case <-pingTimer.C:
			c.logger.Info().Msg("ping")
			requestMessage, err := c.makePingMessage()
			if err != nil {
				c.logger.Info().Msg("ping make error")
				break L
			}
			c.logger.Info().Interface("sendData", requestMessage).Interface("sendBinaryData", binary.Size(requestMessage)).Interface("userID", w.userID).Msg("SEND-PONG-LOG")
			if err := c.sendMessage(&c.conn, requestMessage); err != nil {
				break L
			}
			pongTimer.Reset(10 * time.Second)
		case <-pongTimer.C:
			c.logger.Info().Msg("pong is failed")
			break L
		case msg, ok := <-c.receiveMessage:
			c.logger.Info().Msg(string(msg))
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
			c.logger.Info().Interface("sendData", msg).Interface("sendBinaryData", binary.Size(msg)).Interface("userID", w.userID).Msg("SEND-MESSAGE-LOG")
			// クライアントへメッセージ送信
			if err := c.sendMessage(&c.conn, msg); err != nil {
				break L
			}
		}
	}
	// TODO: サーバが死んだことによる
	// TODO: エラーメッセージをクライアント側へ送る
	if err := c.sendMessage(&c.conn, []byte{}); err != nil {
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

		if length == 0 {
			c.logger.Info().Str("", "d").Msg("connection is closed")
			break L
		} else {
			functionType := data[0]
			print(functionType)
			c.logger.Info().Caller().Msg(string(data))
			c.receiveMessage <- data
		}
		<-ctx.Done()
	}
}

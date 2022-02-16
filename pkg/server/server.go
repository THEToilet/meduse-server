package server

import (
	"context"
	"github.com/rs/zerolog"
	"meduse-server/pkg/server/handler"
	"net"
)

func NewServer(port string, logger *zerolog.Logger) {
	// NOTE: IPv4のみ
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+port)
	if err != nil {
		logger.Fatal().Interface("server down", "d").Msg("")
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logger.Fatal().Interface("server down", "d").Msg("")
	}

	logger.Info().Str("Addr", port).Msg("Serve is running")

	for {

		conn, err := listener.Accept()
		if err != nil {
			logger.Err(err).Msg(err.Error())
			continue
		}

		// NOTE: goroutineのキャンセル処理に使う
		ctx, cancel := context.WithCancel(context.Background())

		receiveMessage := make(chan []byte, 100)
		sendingMessage := make(chan []byte, 100)
		connection := handler.NewConnection(conn, receiveMessage, sendingMessage, logger)
		go connection.Selector(ctx, cancel)
		go connection.Receiver(ctx)

		for {
			data := make([]byte, 128)
			length, err := conn.Read(data)
			if err != nil {
				logger.Err(err).Msg(err.Error())
				break
			}

			if length == 0 {
				logger.Info().Str("", "d").Msg("connection is closed")
				break
			} else {
				functionType := data[0]
				print(functionType)
				n, err := conn.Write([]byte("--"))
				logger.Info().Str("", "d").Msg(string(n))
				if err != nil {
					logger.Err(err).Msg(err.Error())
					break
				}
			}
		}
	}
}

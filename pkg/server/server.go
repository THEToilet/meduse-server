package server

import (
	"context"
	"github.com/rs/zerolog"
	"meduse-server/pkg/domain/application"
	"meduse-server/pkg/server/handler"
	"net"
)

func NewServer(port string, userUseCase *application.UserUseCase, roomUseCase *application.RoomUseCase, logger *zerolog.Logger) {
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
		connection := handler.NewConnection(conn, receiveMessage, sendingMessage, userUseCase, roomUseCase, logger)
		go connection.Selector(ctx, cancel)
		go connection.Receiver(ctx)
	}
}

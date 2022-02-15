package server

import (
	"github.com/rs/zerolog"
	"net"
)

func NewServer(port string, logger *zerolog.Logger) {

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal().Interface("server down", "d").Msg("")
	}
	logger.Info().Str("Addr", port).Msg("Serve is running")

	for{
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatal().Interface("server down", "d").Msg("")
		}
	}
}

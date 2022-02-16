package server

import (
	"github.com/rs/zerolog"
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

package main

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"log"
	"meduse-server/pkg/config"
	"meduse-server/pkg/domain/application"
	"meduse-server/pkg/gateway/repository"
	logger2 "meduse-server/pkg/logger"
	"meduse-server/pkg/server"
	"os"
	"strconv"
)

var (
	version = "0.3.0"
	logger  *zerolog.Logger
	con     *config.Config
)

func init() {
	file, err := os.Open("meduse-server.conf")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	con = config.NewConfig(buffer)
	fmt.Println(con)

	logger, err = logger2.NewLogger(con)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info().Str("Title", con.Title).Msg("Config")
	logger.Info().Str("LogLevel", con.LogInfo.Level).Msg("Config")
}

func main() {

	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("meduse-server version is %s", version)
		return
	}

	// NOTE: Repository
	userRepository := repository.NewUserRepository()
	roomRepository, err := repository.NewRoomRepository()
	if err != nil {
		logger.Fatal().Msg("")
	}

	// NOTE: UseCase
	userUseCase := application.NewUserUseCase(userRepository)
	roomUseCase := application.NewRoomUseCase(roomRepository)

	server.NewServer(strconv.Itoa(int(con.Server.Port)), userUseCase, roomUseCase, logger)

}

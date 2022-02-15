package meduse_server

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"log"
	"meduse-server/pkg/config"
	logger2 "meduse-server/pkg/logger"
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

	logger.Info().Str("Addr", strconv.Itoa(int(con.Server.Port))).Msg("Serve is running")

	/*
		if _, err := os.Stat(con.Cert); err == nil {
			logger.Info().Str("type", "wss").Msg("wss server start")
			if err := server.ListenAndServeTLS(con.Cert, con.Key); err != nil {
				logger.Fatal().Err(err).Msg("wss error")
			}
		} else {
			logger.Info().Str("type", "ws").Msg("ws server start")
			if err := server.ListenAndServe(); err != nil {
				logger.Fatal().Err(err).Msg("ws error")
			}
		}

	*/
}

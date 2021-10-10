package main

import (
	"embed"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"wsTest/configs"
	"wsTest/service"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main()  {
	//reading json file for configs
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}
	//creating config entity to deserialize configs.json
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}
	//echo server
	app := echo.New()
	//channel for errors
	errCh := make(chan error, 1)

	go service.BinanceWSService(cfg, errCh)

	errCh <- app.Start(cfg.Port)
	log.Fatalf("Terminated: %v", <-errCh)
}


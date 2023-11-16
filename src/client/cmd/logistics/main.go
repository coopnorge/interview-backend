package main

import (
	"log"

	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/config"
	"github.com/coopnorge/interview-backend/src/server"
)

func main() {
	cfg := &config.ClientAppConfig{}
	cfg.LoadFromEnv()

	log.Println("Loaded Configuration from Environment Variables\n", cfg)

	app, cleanup, err := newWire(cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if e := app.Run(); e != nil {
		panic(e)
	}

	svc := &coopLogisticsEngineAPIClient{}
	go server.RunGRPCServer(":50051", svc)
}

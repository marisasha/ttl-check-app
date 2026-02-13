package main

import (
	_ "github.com/lib/pq"
	_ "github.com/marisasha/ttl-check-app/docs"
	"github.com/marisasha/ttl-check-app/internal/app"
	"github.com/marisasha/ttl-check-app/internal/config"
	_ "github.com/marisasha/ttl-check-app/internal/docs"
	"github.com/marisasha/ttl-check-app/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err)
	}

	application, err := app.NewApp(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to initialize app: %s", err.Error())
	}

	app.RunWithGracefulShutdown(application, cfg.AppPort)

}

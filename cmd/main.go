package main

import (
	"os"

	_ "github.com/lib/pq"
	ttlchecker "github.com/marisasha/ttl-check-app"
	"github.com/marisasha/ttl-check-app/pkg/handler"
	"github.com/marisasha/ttl-check-app/pkg/repository"
	"github.com/marisasha/ttl-check-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title TTL Checker API
// @version 1.0
// @description API для проверки TTL сертификатов
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@ttl-checker.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Введите: Bearer {token}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error with initializing configs: %s", err.Error())
	}

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error with loading .env: %s", err.Error())
	// }

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(ttlchecker.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

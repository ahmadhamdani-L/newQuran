package main

import (
	db "quran/database"
	"quran/database/migration"
	"quran/internal/factory"
	"quran/internal/http"
	"quran/internal/middleware"
	// "quran/pkg/elasticsearch"
	"quran/pkg/util/env"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title quran
// @version 0.0.1
// @description This is a doc for quran.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv("PORT")

	db.Init()
	migration.Init()
	// elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}

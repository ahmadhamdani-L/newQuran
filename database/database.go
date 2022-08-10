package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	dbConnections map[string]*gorm.DB
)

func Init() {
	dbConfigurations := map[string]Db{
		"QURAN": &dbPostgreSQL{
			db: db{
				Host: os.Getenv("DB_HOST_QURAN"),
				User: os.Getenv("DB_USER_QURAN"),
				Pass: os.Getenv("DB_PASS_QURAN"),
				Port: os.Getenv("DB_PORT_QURAN"),
				Name: os.Getenv("DB_NAME_QURAN"),
			},
			SslMode: os.Getenv("DB_SSLMODE_QURAN"),
			Tz:      os.Getenv("DB_TZ_QURAN"),
		},
	}

	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}

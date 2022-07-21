package database

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"privy/interfaces"
	"privy/pkg/parsing"
)

var config interfaces.SQLConfig

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Error("Unable to find .env file")
		return
	}

	config = interfaces.SQLConfig{
		Enable:   parsing.StrToBoolEnv("DATABASE_ENABLE"),
		Host:     parsing.StrEnv("DATABASE_HOST"),
		Port:     parsing.StrToIntEnv("DATABASE_PORT"),
		Username: parsing.StrEnv("DATABASE_USERNAME"),
		Password: parsing.StrEnv("DATABASE_PASSWORD"),
		Database: parsing.StrEnv("DATABASE_NAME"),
	}
}

func GetConfig() interfaces.SQLConfig {
	return config
}

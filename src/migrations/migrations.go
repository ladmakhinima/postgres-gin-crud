package config

import (
	"github.com/ladmakhinima/postgres-gin-crud/src/config"
	"github.com/ladmakhinima/postgres-gin-crud/src/users"
)

func LoadMigrations() {
	config.DB.AutoMigrate(&users.User{})
}

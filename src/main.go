package main

import (
	gin "github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	configApp "github.com/ladmakhinima/postgres-gin-crud/src/config"
	migrations "github.com/ladmakhinima/postgres-gin-crud/src/migrations"
	users "github.com/ladmakhinima/postgres-gin-crud/src/users"
)

func main() {
	loadEnv()
	server := gin.Default()
	configApp.ConnectDB()
	migrations.LoadMigrations()
	users.LoadUserRoutes(server)
	server.Run()
}

func loadEnv() {
	env.Load("./.env")
}

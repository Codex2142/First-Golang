package main

import (
	"golang-crud-login/config"
	"golang-crud-login/migrations"
	"golang-crud-login/routes"
)

func main() {
	config.ConnectDB()
	migrations.Migrate()
	r := routes.SetupRouter()
	r.Run(":8080")
}

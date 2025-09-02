package migrations

import (
	"golang-crud-login/config"
	"golang-crud-login/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{}, &models.Product{})
}

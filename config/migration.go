package config

import "mini/entity"

func MigrateDB() {
	DB.AutoMigrate(&entity.User{})
}
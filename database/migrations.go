package database

import (
	"log"
)

// TODO: Implement database versioning
func RunMigrations(auto bool) {
	ConnectDb()
	log.Println("Running migrations...")
	if auto {
		DB.Db.AutoMigrate(
			&User{},
			&Schedule{},
			&Speaker{},
			&Comment{},
		)
	}
	log.Println("Done")
}

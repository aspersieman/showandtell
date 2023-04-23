package database

import (
	models "bitbucket.org/envirovisionsolutions/showandtell/models"
	"log"
)

// TODO: Implement migrations
func Migrate(auto bool) {
	ConnectDb()
	log.Println("Running migrations...")
	if auto {
		DB.Db.AutoMigrate(
			&models.User{},
			&models.ScheduleStatus{},
			&models.Schedule{},
			&models.SpeakerStatus{},
			&models.Speaker{},
			&models.Comment{},
		)
	}
	log.Println("Done")
}

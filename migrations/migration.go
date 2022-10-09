package migrations

import (
	"gin-boilerplate/infra/database"
	"gin-boilerplate/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var exampleModels = []interface{}{&models.Example{}}
	exampleModels = append(exampleModels, &models.User{})
	err := database.DB.AutoMigrate(exampleModels...)
	if err != nil {
		return
	}
}

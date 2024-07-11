package migrate

import (
	"log"

	"github.com/somtojf/gosh/initializers"
	"github.com/somtojf/gosh/models"
)

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.History{})

	if err != nil {
		log.Fatal(err.Error())
	}
}

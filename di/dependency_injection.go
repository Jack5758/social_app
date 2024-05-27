package di

import (
	"github.com/Jack5758/infra/db"
	"github.com/Jack5758/internal/config"
	"log"
)

func RunApp() error {

	appConfig, err := config.NewAppConfig()
	if err != nil {
		return err
	}
	_, err = db.SetupDB(appConfig.DB)
	if err != nil {
		return err
	}

	log.Println("successfully connected to DB")
	return nil
}

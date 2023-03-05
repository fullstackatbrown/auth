package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	// Default config has 10 second connection timeout
	if err := mgm.SetDefaultConfig(nil, "auth", options.Client().ApplyURI(config.Config.MongoUri)); err != nil {
		panic(err)
	}
}

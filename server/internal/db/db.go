package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	// Default config has 10 second connection timeout
	if err := mgm.SetDefaultConfig(nil, config.Config.DbName, options.Client().ApplyURI(config.Config.MongoUri)); err != nil {
		panic(err)
	}
}

func Create(model mgm.Model) (err error) {
	err = mgm.Coll(model).Create(model)
	return
}

func Update(model mgm.Model, upsert bool) (err error) {
	err = mgm.Coll(model).Update(model, &options.UpdateOptions{Upsert: &upsert})
	return
}

func Delete(model mgm.Model) (err error) {
	err = mgm.Coll(model).Delete(model)
	return
}

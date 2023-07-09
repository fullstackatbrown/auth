package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func FindDomainByName(name string) (domain *model.Domain, err error) {
	domain = &model.Domain{}
	err = mgm.Coll(domain).FindOne(mgm.Ctx(), bson.M{"name": name}).Decode(domain)
	return
}

func DeleteDomain(name string) (err error) {
	domain, err := FindDomainByName(name)
	if err != nil {
		return err
	}
	err = Delete(domain)
	return
}

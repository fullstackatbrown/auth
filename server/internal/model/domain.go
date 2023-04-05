package model

import "github.com/kamva/mgm/v3"

type Domain struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Roles            Roles  `json:"roles" bson:"roles"`
}

func NewDomain(name string, roles Roles) *Domain {
	return &Domain{
		Name:  name,
		Roles: roles,
	}
}

func (model *Domain) CollectionName() string {
	return "domains"
}

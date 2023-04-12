package model

import "github.com/kamva/mgm/v3"

type Domain struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Roles            []Role `json:"roles,omitempty" bson:"roles,omitempty"`
}

func (model *Domain) CollectionName() string {
	return "domains"
}

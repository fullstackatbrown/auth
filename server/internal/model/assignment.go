package model

import "github.com/kamva/mgm/v3"

type Assignment struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	Role             Role   `json:"roles" bson:"role"`
}

func (model *Assignment) CollectionName() string {
	return "assignments"
}

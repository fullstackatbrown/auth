package model

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pronouns         string `json:"pronouns" bson:"pronouns"`
}

func NewUser(name string, pronouns string) *User {
	return &User{
		Name:     name,
		Pronouns: pronouns,
	}
}

func (model *User) CollectionName() string {
	return "users"
}

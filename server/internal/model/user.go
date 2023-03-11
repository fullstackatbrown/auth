package model

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Pronouns         string `json:"pronouns,omitempty" bson:"pronouns,omitempty"`
}

func NewUser(name string, email string, pronouns string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Pronouns: pronouns,
	}
}

func (model *User) CollectionName() string {
	return "users"
}

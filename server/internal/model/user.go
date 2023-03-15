package model

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Profile          Profile `json:"profile" bson:"profile"`
	Roles            Roles   `json:"roles,omitempty" bson:"roles,omitempty"`
}

type Profile struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email" bson:"email"`
	Pronouns string `json:"pronouns,omitempty" bson:"pronouns,omitempty"`
}

type Roles []string

func NewUser(email string) *User {
	return &User{
		Profile: Profile{
			Email: email,
		},
	}
}

func (model *User) CollectionName() string {
	return "users"
}

package model

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	GoogleId         string  `json:"googleId" bson:"googleId"`
	Profile          Profile `json:"profile" bson:"profile"`
	Roles            Roles   `json:"roles" bson:"roles"`
}

type Profile struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Pronouns string `json:"pronouns,omitempty" bson:"pronouns,omitempty"`
}

func NewUser(googleId string, name string, email string) *User {
	return &User{
		GoogleId: googleId,
		Profile: Profile{
			Name:  name,
			Email: email,
		},
		Roles: []Role{},
	}
}

func (model *User) CollectionName() string {
	return "users"
}

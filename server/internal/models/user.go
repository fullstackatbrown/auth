package models

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"firstName" bson:"firstName"`
	LastName         string `json:"lasttName" bson:"lastName"`
	Pronouns         string `json:"pronouns" bson:"pronouns"`
}

func NewUser(firstName string, lastName string, pronouns string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Pronouns:  pronouns,
	}
}

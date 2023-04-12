package model

type Role struct {
	Domain string `json:"domain,omitempty" bson:"domain,omitempty"`
	Role   string `json:"role" bson:"role"`
}

package model

type Roles []Role

type Role struct {
	Domain string `json:"domain" bson:"domain"`
	Role   string `json:"role" bson:"role"`
}

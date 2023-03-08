package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/kamva/mgm/v3"
)

func CreateUser(user *model.User) (err error) {
	err = mgm.Coll(user).Create(user)
	return
}

func FindUserById(id string) (user *model.User, err error) {
	user = &model.User{}
	err = mgm.Coll(user).FindByID(id, user)
	return
}

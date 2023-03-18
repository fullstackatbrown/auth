package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func FindUsersByEmail(email string) (users []model.User, err error) {
	users = []model.User{}
	err = mgm.Coll(&model.User{}).SimpleFind(&users, bson.M{"profile.email": email})
	return
}

func UpsertUser(user *model.User) (err error) {
	upsert := true
	err = mgm.Coll(user).Update(user, &options.UpdateOptions{Upsert: &upsert})
	return
}

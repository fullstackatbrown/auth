package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
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

func FindUserByEmail(email string) (user *model.User, err error) {
	users := []model.User{}
	err = mgm.Coll(user).SimpleFind(&users, bson.M{"email": bson.M{operator.Eq: email}})
	user = &users[0]
	return
}

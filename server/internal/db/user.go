package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUserById(id string) (user *model.User, err error) {
	user = &model.User{}
	err = mgm.Coll(user).FindOne(mgm.Ctx(), bson.M{"googleId": id}).Decode(user)
	return
}

func FindUsersByEmail(email string) (users []model.User, err error) {
	users = []model.User{}
	err = mgm.Coll(&model.User{}).SimpleFind(&users, bson.M{"profile.email": email})
	return
}

func DeleteUser(userId string) (err error) {
	// find this user first
	user, err := FindUserById(userId)
	if err != nil {
		return err
	}
	err = Delete(user)
	return
}

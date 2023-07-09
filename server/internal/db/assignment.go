package db

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAssignmentById(id string) (assignment *model.Assignment, err error) {
	assignment = &model.Assignment{}
	err = mgm.Coll(assignment).FindOne(mgm.Ctx(), bson.M{"_id": id}).Decode(assignment)
	return
}

func FindAssignmentsByEmail(email string) (assignments []model.Assignment, err error) {
	assignments = []model.Assignment{}
	err = mgm.Coll(&model.Assignment{}).SimpleFind(&assignments, bson.M{"email": email})
	return
}

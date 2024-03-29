package handler

import (
	"net/http"

	"github.com/fullstackatbrown/auth-infrastructure/internal/db"
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsersByEmail(w http.ResponseWriter, r *http.Request) {
	// get email from query param
	email := r.URL.Query().Get("email")

	// handle error
	users, _ := db.FindUsersByEmail(email)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// delete user from db
	err := db.DeleteUser(userId)
	if err != nil && err != mongo.ErrNoDocuments {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, map[string]string{"message": "user deleted"})
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// get user object from db
	user, err := db.FindUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "user not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// get updated user object from request body
	err = render.DecodeJSON(r.Body, &user.Profile)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid request body"})
		return
	}

	// persist update to db
	err = db.Update(user, false)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusOK)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// get user object from db
	user, err := db.FindUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "user not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// get profile from user object
	profile := user.Profile

	render.Status(r, http.StatusOK)
	render.JSON(w, r, profile)
}

func ListUserRoles(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// get user object from db
	user, err := db.FindUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "user not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// get roles from user object
	roles := user.Roles

	render.Status(r, http.StatusOK)
	render.JSON(w, r, roles)
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// get user object from db
	user, err := db.FindUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "user not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// get newRole from body
	var newRole model.Role
	err = render.DecodeJSON(r.Body, &newRole)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid request body"})
		return
	}

	// check that user doesn't already have role
	for _, role := range user.Roles {
		if role.Domain == newRole.Domain && role.Role == newRole.Role {
			render.Status(r, http.StatusOK)
			return
		}
	}

	// add role to user object if there isn't one already
	add := true
	for _, role := range user.Roles {
		if role.Domain == newRole.Domain && role.Role == newRole.Role {
			add = false
			break
		}
	}
	if add {
		user.Roles = append(user.Roles, newRole)
	}

	// persist update to db
	err = db.Update(user, false)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusCreated)
}

func RemoveUserRole(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// get user object from db
	user, err := db.FindUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "user not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// get domain and role from query params
	domain := r.URL.Query().Get("domain")
	role := r.URL.Query().Get("role")

	// remove corresponding role from user object
	for i, r := range user.Roles {
		if r.Domain == domain && r.Role == role {
			user.Roles = append(user.Roles[:i], user.Roles[i+1:]...)
			break
		}
	}

	// persist update to db
	err = db.Update(user, false)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusNoContent)
}

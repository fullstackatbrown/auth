package handler

import (
	"net/http"

	"github.com/fullstackatbrown/auth-infrastructure/internal/db"
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	// decode the request body into a new Domain
	newDomain := &model.Domain{}
	err := render.DecodeJSON(r.Body, newDomain)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid request body"})
		return
	}

	// check if the domain already exists
	_, err = db.FindDomainByName(newDomain.Name)
	if err != mongo.ErrNoDocuments {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, map[string]string{"message": "domain already exists"})
		return
	}

	// insert the new Domain into the db
	err = db.Update(newDomain, true)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// return the new Domain
	render.Status(r, http.StatusCreated)
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	// get domain from path param
	name := chi.URLParam(r, "domainName")

	// get domain from db
	domain, err := db.FindDomainByName(name)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// return domain
	render.Status(r, http.StatusOK)
	render.JSON(w, r, domain)
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	// get domain from path param
	name := chi.URLParam(r, "domainName")

	// get domain from db
	err := db.DeleteDomain(name)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// return success
	render.Status(r, http.StatusNoContent)
}

func CreateDomainRole(w http.ResponseWriter, r *http.Request) {
	// get domain from path param
	name := chi.URLParam(r, "domainName")

	// get domain from db
	domain, err := db.FindDomainByName(name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "domain not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// decode the request body into a new Role
	newRole := model.Role{}
	err = render.DecodeJSON(r.Body, &newRole)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid request body"})
		return
	}

	// add role to domain if there isn't one already
	add := true
	for _, r := range domain.Roles {
		if r.Role == newRole.Role {
			add = false
			break
		}
	}
	if add {
		domain.Roles = append(domain.Roles, newRole)
	}

	// update domain in db
	err = db.Update(domain, false)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// return success
	render.Status(r, http.StatusCreated)
}

func DeleteDomainRole(w http.ResponseWriter, r *http.Request) {
	// get domain from path param
	name := chi.URLParam(r, "domainName")

	// get role from query param
	role := r.URL.Query().Get("role")

	// get domain from db
	domain, err := db.FindDomainByName(name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "domain not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// remove role from domain
	for i, r := range domain.Roles {
		if r.Role == role {
			domain.Roles = append(domain.Roles[:i], domain.Roles[i+1:]...)
			break
		}
	}

	// update domain in db
	err = db.Update(domain, false)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// return success
	render.Status(r, http.StatusNoContent)
}

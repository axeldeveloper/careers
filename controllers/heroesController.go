package controllers

import (
	"encoding/json"
	"levpay/models"
	u "levpay/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateHeroe(w http.ResponseWriter, r *http.Request) {
	hero := &models.Heroe{}

	err := json.NewDecoder(r.Body).Decode(hero)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := hero.CreateHeroe()
	u.Respond(w, resp)
}

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroe := &models.Heroe{}
	data := heroe.GetHeroes()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetHeroe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	heroe := &models.Heroe{}
	data := heroe.GetHeroe(pid)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func UpdateHeroe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}
	hero := &models.Heroe{}
	err = json.NewDecoder(r.Body).Decode(hero)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data, err := hero.UpdateHeroe(uid)
	if err != nil {
		u.Respond(w, u.Message(false, "Error Update"))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func DeleteHeroe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	heroe := &models.Heroe{}
	data, err := heroe.DeleteHeroe(pid)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

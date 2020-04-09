package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"levpay/models"
	u "levpay/utils"

	"net/http"
	"strconv"
)

func CreateHeroe(w http.ResponseWriter, r *http.Request) {
	hero := &models.Heroe{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &hero)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, err.Error()))
		return
	}
	resp := hero.CreateHeroe()

	if resp["status"].(bool) == false {
		u.Respond(w, http.StatusBadRequest, resp)
		return
	}
	u.Respond(w, http.StatusCreated, map[string]interface{}{"id": resp["ID"]})
}

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroe := &models.Heroe{}
	data, err := heroe.GetHeroes()
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Erro list super heroes"))
		//return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

func GetHeroe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid := vars["id"]
	heroe := &models.Heroe{}
	data, err := heroe.GetHeroe(pid)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Erro get super heroes"))
		//return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

func GetNameHeroe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	heroe := &models.Heroe{}
	data, err := heroe.GetNameHeroe(name)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
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
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	data, err := hero.UpdateHeroe(uid)
	if err != nil {
		//u.Respond(w, u.Message(false, "Error Update"))
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusAccepted, resp)
}

func DeleteHeroe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//user := r.Context().Value("userID").(float64)
	//userID := uint(user)

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
	u.Respond(w, http.StatusAccepted, resp)
}

func TesteHttp(w http.ResponseWriter, r *http.Request) {
	//var err error
	//err = "Error ERROR"
	resp := u.Message(true, "success")
	resp["data"] = "teste"
	//u.ERROR(w, http.StatusBadRequest, err)
	//u.Respond(w, http.StatusCreated, resp)
	u.Respond(w, http.StatusAccepted, resp)
}

/*
func GetIntelligenceHeroe(w http.ResponseWriter, r *http.Request) {
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
	u.Respond(w, http.StatusOK, resp)
}
*/

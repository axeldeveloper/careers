package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"levpay/models"
	u "levpay/utils"

	"github.com/gorilla/mux"

	"net/http"
)

func getType(r *http.Request) string {
	//'Hero',
	//'Villain'
	vars := mux.Vars(r)
	param := vars["type"]
	fmt.Println(r.URL)
	fmt.Println("Tipo get ", param)
	var value string
	if param == "" {
		value = "Todo"
	} else {
		value = param
	}

	return value
}

// @todo get all supers
func GetSupers(w http.ResponseWriter, r *http.Request) {
	heroe := &models.Super{}
	data, err := heroe.GetSupers()
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Erro list super heroes"))
		//return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

// @todo get super by type
func GetSupersType(w http.ResponseWriter, r *http.Request) {
	super := &models.Super{}

	param := getType(r)
	fmt.Println("par", param)
	data, err := super.GetSupersType(param)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Erro list super heroes"))
		//return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

// @todo get super by ID
func GetSuper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid := vars["id"]

	heroe := &models.Super{}
	data, err := heroe.GetSuper(pid)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, "Erro get super heroes"))
		//return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

// @todo super get by name
func GetNameSuper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	super := &models.Super{}
	data, err := super.GetNameSuper(name)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

// @todo create super
func CreateSuper(w http.ResponseWriter, r *http.Request) {
	super := &models.Super{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &super)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, err.Error()))
		return
	}

	if resp, ok := super.Validate(); !ok {
		u.Respond(w, http.StatusBadRequest, resp)
		return
	}

	data, err := super.CreateSuper()
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, err.Error()))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	//u.JSON(w, http.StatusCreated, data.ID)
	u.Respond(w, http.StatusCreated, resp)
}

// @todo update super
func UpdateSuper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	super := &models.Super{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &super)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.Message(false, err.Error()))
		return
	}

	data, err := super.UpdateSuper(uid)

	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if data != nil {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, http.StatusNoContent, resp)
	} else {
		u.ERROR(w, http.StatusBadRequest, err)
	}

}

// @todo delete super
func DeleteSuper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid := vars["id"]
	super := &models.Super{}
	data, err := super.DeleteSuper(pid)
	if err != nil {
		u.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusNoContent, resp)
}

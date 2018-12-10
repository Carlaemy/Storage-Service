package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../../DB/Logica"
	"../Model"
	"github.com/gorilla/mux"
)

var users []Model.User

// GetUser...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users = data.Read_User_DB()
	respondJSON(w, http.StatusOK, users)
}

// CreateUser...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := Model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	data.Write_User_DB(user, false)
	respondJSON(w, http.StatusCreated, user)
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, user := range data.Read_User_DB() {
		if user.UserID == id {
			respondJSON(w, http.StatusOK, user)
			return
		}
	}
	respondJSON(w, http.StatusNotFound, nil)
}

// DeleteUser...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, user := range data.Read_User_DB() {
		if user.UserID == id {
			data.Write_User_DB(user, true)
			respondJSON(w, http.StatusOK, "Delected successful")
			return
		}
	}
	respondJSON(w, http.StatusNotFound, nil)
}

// Login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	login := Model.Login{}

	err := r.ParseForm()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&login); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	status := data.Login(login.Email, login.Password)
	if status == "Success" {
		respondJSON(w, http.StatusOK, status)
		return
	}
	respondJSON(w, http.StatusNotFound, status)
}

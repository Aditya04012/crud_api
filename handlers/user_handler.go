package handlers

import (
	"database/sql"
	"encoding/json"
	"golang-server/golang_api/models"
	"golang-server/golang_api/repository"
	"net/http"
	"strconv"
)

type UserHandler struct{
	DB *sql.DB
}

func (h *UserHandler)GetUser(w http.ResponseWriter,r *http.Request){
	users,err:=repository.GetUsers(h.DB)
	if err!=nil{
		http.Error(w,"Failed  to fetchUser",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","Application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler)DeleteUser(w http.ResponseWriter,r *http.Request){
	str:=r.PathValue("id")
	id,err:=strconv.Atoi(str)

	if err!=nil{
		http.Error(w,"Invalid User id",http.StatusBadRequest)
		return
	}

	_,err1:=repository.DeleteUserById(h.DB,id)
	if err1!=nil{
		http.Error(w,"User Not Found",http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler)CreateUser(w http.ResponseWriter,r *http.Request){
	var user models.User

	err:=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
	}

	if user.Name=="" || user.Email==""{
		http.Error(w,"Name and email is required",http.StatusBadRequest)
		return
	}

    err=repository.Create(h.DB,&user)

	if err!=nil{
		http.Error(w, "Failed to create user",http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
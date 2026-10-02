package main

import (
	"fmt"
	"golang-server/golang_api/db"
	"net/http"
	"golang-server/golang_api/handlers"
)

func main(){
  database:=db.Connect()
  defer database.Close()

  mux:=http.NewServeMux()
 
  UserHandler:=&handlers.UserHandler{
    DB:database,
  }

  mux.HandleFunc("GET /users",UserHandler.GetUser)
  mux.HandleFunc("POST /users",UserHandler.CreateUser)
  mux.HandleFunc("DELETE /users/{id}",UserHandler.DeleteUser)

  server:=http.Server{
	Addr: ":3000",
	Handler: mux,
  }

  fmt.Println("Server is live at port 3000")
  server.ListenAndServe()
}
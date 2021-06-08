package cmd

import (
	"log"
	"net/http"

	//"Assignemnts/APIs/controller"

	"github.com/gorilla/mux"
	"github.com/lucky-786/gorm-api/controller"
)

func Execute() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/v1/user/{id}", controller.GetSingleUser).Methods("GET")
	myRouter.HandleFunc("/api/v1/user/fetch", controller.GetMultiUser).Methods("POST")
	myRouter.HandleFunc("/api/v1/user/create", controller.CreateUser).Methods("POST")
	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

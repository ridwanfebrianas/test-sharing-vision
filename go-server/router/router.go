package router

import (
	"test-sharing-vision/go-server/middleware"

	"github.com/gorilla/mux"
)

// Router . . .
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/user", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/{limit}/{offset}", middleware.GetAllUserPagination).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/{id}", middleware.GetUserByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/user/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
	return router
}

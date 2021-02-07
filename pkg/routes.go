package pkg

import (
	"github.com/gorilla/mux"
	"github.com/mfsiddique11/auth-server/db"
	"github.com/mfsiddique11/auth-server/pkg/user"
)

func Routes() *mux.Router {

	mongoClient := user.MongoClient{Database: db.InitMongoDB()}

	router := mux.NewRouter()

	// user's routes
	router.HandleFunc("/v1/users", mongoClient.CreateUsersHandler).Methods("POST")
	router.HandleFunc("/v1/users", mongoClient.GetUsersHandler).Methods("GET")
	router.HandleFunc("/v1/users/{id}", mongoClient.GetUserHandler).Methods("GET")
	router.HandleFunc("/v1/users/{id}", mongoClient.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/v1/users/{id}", mongoClient.DeleteUserHandler).Methods("DELETE")

	return router
}

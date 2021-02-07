package user

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os/user"
	"time"
)

type MongoClient struct {
	Database *mongo.Database
}

func (client MongoClient) CreateUsersHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var u user.User

	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err := client.Database.Collection("users").InsertOne(ctx, u)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(res)
}

func (client MongoClient) GetUsersHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := client.Database.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(users)

}

func (client MongoClient) GetUserHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println("Invalid id")
	}

	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := client.Database.Collection("users").Find(ctx, bson.M{"_id": userId})
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(cursor)

}

func (client MongoClient) UpdateUserHandler(response http.ResponseWriter, request *http.Request) {
	var u user.User

	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	params := mux.Vars(request)
	userId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println("Invalid id")
	}

	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := client.Database.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": userId},
		bson.D{
			{"$set", u},
		})
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(cursor)

}

func (client MongoClient) DeleteUserHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println("Invalid id")
	}

	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err := client.Database.Collection("users").DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedCount == 0 {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	response.WriteHeader(http.StatusNoContent)

}

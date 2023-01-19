package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/mongodb/mongo-go-driver"
	"github.com/julieschmidt/httprouter"
)
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, _ = mongo.Connect(ctx, clientOptions)

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}
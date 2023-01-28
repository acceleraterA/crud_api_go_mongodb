package main

import (
	"net/http"

	"github.com/acceleraterA/crud_api_go_mongodb/controllers"
	"github.com/gorilla/mux"
)

// type Person struct {
// 	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
// 	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
// }

// // var client *mongo.Client
// var collection *mongo.Collection

//	func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
//		//declare the response type json
//		fmt.Fprintf(response, "Application is up and running \n")
//		response.Header().Set("content-type", "application/json")
//		var person Person
//		json.NewDecoder(request.Body).Decode(&person)
//		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//		result, _ := collection.InsertOne(ctx, person)
//		json.NewEncoder(response).Encode(result)
//	}
//
//	func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
//		response.Header().Set("content-type", "application/json")
//		params := mux.Vars(request)
//		id, _ := primitive.ObjectIDFromHex(params["id"])
//		var person Person
//		//collection := client.Database("thepolyglotdeveloper").Collection("people")
//		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//		err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
//		if err != nil {
//			response.WriteHeader(http.StatusInternalServerError)
//			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
//			return
//		}
//		json.NewEncoder(response).Encode(person)
//	}
//
//	func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
//		response.Header().Set("content-type", "application/json")
//		var people []Person
//		//collection := client.Database("thepolyglotdeveloper").Collection("people")
//		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//		cursor, err := collection.Find(ctx, bson.M{})
//		if err != nil {
//			response.WriteHeader(http.StatusInternalServerError)
//			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
//			return
//		}
//		defer cursor.Close(ctx)
//		for cursor.Next(ctx) {
//			var person Person
//			cursor.Decode(&person)
//			people = append(people, person)
//		}
//		if err := cursor.Err(); err != nil {
//			response.WriteHeader(http.StatusInternalServerError)
//			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
//			return
//		}
//		json.NewEncoder(response).Encode(people)
//	}
func main() {
	// fmt.Println("Starting the application...")
	// // load .env file
	// err := godotenv.Load("../.env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	// clientOptions := options.Client().
	// 	ApplyURI(os.Getenv("DB_URI")).
	// 	SetServerAPIOptions(serverAPIOptions)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Check the connection
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

	// collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION_NAME"))

	// fmt.Println("Collection instance created!")
	router := mux.NewRouter()
	router.HandleFunc("/person", controllers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", controllers.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":3000", router)
}

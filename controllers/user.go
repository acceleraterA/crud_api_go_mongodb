package controllers
import(
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	models "github.com/acceleraterA/crud_api_go_mongodb/models"
)
// var client *mongo.Client
var collection *mongo.Collection
// create connection with mongo db
func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
fmt.Println("Starting the application...")
// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// //clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
// if err != nil {
// 	log.Fatal(err)
// }
// load .env file
err := godotenv.Load("../.env")

if err != nil {
	log.Fatalf("Error loading .env file")
}
serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
clientOptions := options.Client().
	ApplyURI(os.Getenv("DB_URI")).
	SetServerAPIOptions(serverAPIOptions)
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
	log.Fatal(err)
}
// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
	log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")


collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION_NAME"))

fmt.Println("Collection instance created!")


func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	//declare the response type json
	fmt.Fprintf(response, "Application is up and running \n")
	response.Header().Set("content-type", "application/json")
	var person models.Person
	json.NewDecoder(request.Body).Decode(&person)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person models.Person
	//collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people models.[]Person
	//collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}
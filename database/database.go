package database

import (
	"SamirGG/Tesg-Golang/models"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

func getDotEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

// Obtener la URI de la base de datos
// Input: URI del .env
// Output: URI de la base de datos
func getMongoURI() string {
	uri := getDotEnv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set")
	}
	return uri
}

// Conectar a la base de datos
// Input: URI de la base de datos
// Output: cliente
func ConnectToDatabase() error {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		return err
	}

	// Verificar la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	database = client.Database("eventos")
	return nil
}

// Obtener la base de datos
// Input: cliente
// Output: base de datos
func GetDatabase(client *mongo.Client, name string) *mongo.Database {
	return client.Database(name)
}

// Obtener la colección
// Input: base de datos
// Output: colección
func GetCollection(database *mongo.Database, name string) *mongo.Collection {
	return database.Collection(name)
}

// Obtener la colección de eventos
// Input: base de datos
// Output: colección de eventos
func GetEventCollection(database *mongo.Database) *mongo.Collection {
	return GetCollection(database, "events")
}

// Desconectar de la base de datos
// Trae la base de datos y la desconecta
func DisconnectFromDatabase() error {
	return client.Disconnect(context.Background())
}

// Insertar un evento
// Input: evento
// Output: evento insertado
func InsertEvent(event models.Event) (*mongo.InsertOneResult, error) {
	collection := database.Collection("events")
	return collection.InsertOne(context.Background(), event)
}

// Buscar eventos
// Input: evento
// Output: evento encontrado
func FindEvents() ([]models.Event, error) {
	collection := database.Collection("events")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var events []models.Event
	if err = cursor.All(context.Background(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

// Buscar evento por ID
// Input: evento
// Output: evento encontrado
func FindEventById(id primitive.ObjectID) (*models.Event, error) {
	collection := database.Collection("events")
	var event models.Event
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Actualizar un evento
// Input: evento
// Output: evento actualizado
func UpdateEvent(event models.Event) (*mongo.UpdateResult, error) {
	collection := database.Collection("events")
	filter := bson.M{"_id": event.ID}
	update := bson.M{"$set": event}
	return collection.UpdateOne(context.Background(), filter, update)
}

// Eliminar un evento
// Input: evento
// Output: evento eliminado
func DeleteEvent(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := database.Collection("events")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// GetClient retorna el cliente de MongoDB para pruebas
func GetClient() *mongo.Client {
	return client
}

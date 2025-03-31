package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Obtener los datos del .Evn
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
func ConnectToDatabase() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		return nil, err
	}
	return client, nil
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
func DisconnectFromDatabase(client *mongo.Client) error {
	return client.Disconnect(context.Background())
}

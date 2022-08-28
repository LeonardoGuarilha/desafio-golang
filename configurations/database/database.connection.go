package database

import (
	"context"
	"fmt"
	"github.com/leonardoguarilha/configurations"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func ConnectDataBase() *mongo.Collection {
	ctx := context.TODO()

	configuration := configurations.GetConfiguration()
	mongoConnection := options.Client().ApplyURI(configuration.Mongo.Server)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)

	if err != nil {
		log.Fatal("erro ao tentar conexão com o mongo", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("erro ao tentar ping no mongo", err)
	}

	fmt.Println("conexão com o mongo estabelecida com sucesso!")

	return mongoClient.Database(configuration.Mongo.Database).Collection(configuration.Mongo.Collection)
}

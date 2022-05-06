package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/explabs/ad-ctf-paas-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var collection, flags, scoreboard, configurations *mongo.Collection

var ctx = context.TODO()

func InitMongo() {
	adminPass := os.Getenv("ADMIN_PASS")
	if adminPass == "" {
		adminPass = "admin"
	}
	credential := options.Credential{
		Username: "admin",
		Password: adminPass,
	}

	mongoAddr := os.Getenv("MONGODB")
	if mongoAddr == "" {
		mongoAddr = "localhost:27017"
	}
	mongoURI := fmt.Sprintf("mongodb://%s", mongoAddr)
	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("ad").Collection("teams")
	flags = client.Database("ad").Collection("flags")
	scoreboard = client.Database("ad").Collection("scoreboard")
	configurations = client.Database("ad").Collection("config")
}

func GetAuthTeam(login string) (models.Team, error) {
	var team models.Team
	err := collection.FindOne(ctx, bson.M{"login": login}).Decode(&team)
	return team, err
}

package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//db table list
const (
	userTable = "user"
	gistTable = "gist"
)

//DB is db type
type DB struct {
	User *mongo.Collection
	Gist *mongo.Collection
}

//GetDBInstance makes an instance of DB
func getDBInstance() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(CFG.MongoInfo.URL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	currDB := client.Database(CFG.MongoInfo.DBName)

	return &DB{
		User: currDB.Collection(userTable),
		Gist: currDB.Collection(gistTable),
	}

}

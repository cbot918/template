package main

import (
	"context"
	"fmt"
	"os"
	"github.com/cbot918/go-mongo-crud/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri []byte

const (
	port = ":3002"
	dbname = "testdb"
	coname = "testco"
)

func main(){
	var err error
	mongoUri,err = os.ReadFile("mongouri"); if err != nil { panic(err) }

	fmt.Println(string(mongoUri))

	// connect mongodb
	mclient := connectMongodb(string(mongoUri))

	// api init
	api := api.InitApi(mclient)
	
	// api listen
	api.Listen(port)
	
}

func connectMongodb(mongoUri string) *mongo.Client{
	// mongo connect
	
	mclient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri));
	if err != nil{ panic(err) }
	fmt.Printf("\nmongo connect success\n\n")
	return mclient
}
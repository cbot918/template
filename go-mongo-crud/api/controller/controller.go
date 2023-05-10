package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Controller struct {
	mClient *mongo.Client
	collection *mongo.Collection
}
func New(mClient *mongo.Client) *Controller{

		// connect db+collection
	collection := mClient.Database("testdb").Collection("testco")

	return &Controller{
		mClient: mClient,
		collection: collection,
	}
}

func (c *Controller) Ping(ctx *fiber.Ctx) error {
	if err := c.mClient.Ping(context.TODO(), readpref.Primary()); err != nil { panic(err) }
	fmt.Println("request_in: ping mongo")
	return ctx.SendString("ping mongo success")
}

func (c *Controller) UseCollection(ctx *fiber.Ctx) error {
	// WIP
	dbName := ctx.Params("dbname")
	coName := ctx.Params("coname")
	fmt.Printf("dbname: %s\nconame: %s\n\n",dbName, coName)
	fmt.Println("api使用小提醒: 建議預先在資料庫建好db跟collection")

	c.collection = c.mClient.Database(dbName).Collection(coName)

	fmt.Println("request_in: use collection")

	return ctx.SendString("use collection success")
}



func (c *Controller) Insert(ctx *fiber.Ctx) error {
	// do request
	data := &struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}{}
	if err := ctx.BodyParser(data); err != nil { panic(err) }
	fmt.Printf("name: %s, job: %s\n",data.Name, data.Job)

	// do insert mongo
	user := bson.D{{"name", data.Name}, {"age", data.Job}}
	fmt.Printf("\n\ntype: %T\n\n", user)
	result, err := c.collection.InsertOne(context.TODO(), user); if err != nil{ panic(err) }
	fmt.Printf("\n insert success, id: %s \n\n",result.InsertedID)

	// do respond
	res, err := json.Marshal(data); if err != nil{ panic(err) }
	return ctx.Send(res)
}




// 	// insert many
// 	users := []interface{}{
// 		bson.D{{"fullName", "User 2"}, {"age", 25}},
// 		bson.D{{"fullName", "User 3"}, {"age", 20}},
// 		bson.D{{"fullName", "User 4"}, {"age", 28}},
// 	}
// 	results, err := testCollection.InsertMany(context.TODO(), users);if err != nil { panic(err) }
// 	fmt.Println(results.InsertedIDs)
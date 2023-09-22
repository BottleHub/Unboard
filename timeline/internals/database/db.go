package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bottlehub/unboard/boards/configs"
	"github.com/bottlehub/unboard/boards/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return &DB{client: client}, err
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("UserBase").Collection(collectionName)
}

func (db *DB) ctxDeferHelper(collectionName string) (*mongo.Collection, context.Context) {
	collection := colHelper(db, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println(cancel)

	return collection, ctx
}

func (db *DB) resErrHelper(collectionName string, input any) (*mongo.InsertOneResult, error) {
	collection, ctx := db.ctxDeferHelper(collectionName)

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func (db *DB) multipleFetchHelper(collectionName string, ID string, IDName string) (*mongo.Cursor, context.Context) {
	collection, ctx := db.ctxDeferHelper(collectionName)

	objId, _ := primitive.ObjectIDFromHex(ID)

	res, err := collection.Find(ctx, bson.M{IDName: objId})
	if err != nil {
		log.Fatal(err)
	}

	return res, ctx
}

func (db *DB) deleteHelper(collectionName string, ID string) error {
	collection, ctx := db.ctxDeferHelper("messages")

	objId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objId}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}

	return err
}

func (db *DB) updateHelper(collectionName, ID string, info bson.M, model any) error {
	collection, ctx := db.ctxDeferHelper(collectionName)
	_id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": info}

	results := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))
	if err := results.Decode(model); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (db *DB) CreateChatboard(input *model.NewChatboard) (*model.Chatboard, error) {
	res, err := db.resErrHelper("chatboards", input)

	chatboard := &model.Chatboard{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}

	return chatboard, err
}

func (db *DB) GetTimelines(ID string) ([]*model.Timeline, error) {
	res, ctx := db.multipleFetchHelper("timeline", ID, "parentID")
	var (
		timeline []*model.Timeline
		err      error
	)

	defer res.Close(ctx)

	if err = res.All(context.TODO(), timeline); err != nil {
		panic(err)
	}

	return timeline, err
}

func (db *DB) SingleTimeline(ID string) (*model.Timeline, error) {
	collection, ctx := db.ctxDeferHelper("chatboards")
	var timeline *model.Timeline

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&timeline)

	return timeline, err
}

func (db *DB) UpdateTimeline(ID string, input *model.UpdateTimeline) (*model.Timeline, error) {
	var message *model.Timeline

	updateInfo := bson.M{}

	if input.ImageURL != nil {
		updateInfo["imageURL"] = input.ImageURL
	}
	if input.Name != nil {
		updateInfo["name"] = input.Name
	}
	if input.SubTimeline != nil {
		timeline1, err := db.SingleChatboard(ID)
		timeline2, err := db.SingleChatboard(*input.SubTimeline)
		if err != nil {
			panic(err)
		}

		arr := timeline1.SubTimelines
		arr = append(arr, timeline2)
		updateInfo["subTimeline"] = arr
	}
	if input.Text != nil {
		updateInfo["text"] = input.Text
	}

	err := db.updateHelper("timeline", ID, updateInfo, message)

	return message, err
}

func (db *DB) DeleteTimeline(ID string) (*model.DeleteTimeline, error) {
	err := db.deleteHelper("timeline", ID)
	var delete *model.DeleteTimeline

	return delete, err
}

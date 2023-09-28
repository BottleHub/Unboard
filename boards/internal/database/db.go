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

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ConnectDB() (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
	collection, ctx := db.ctxDeferHelper(collectionName)

	objId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objId}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}

	return err
}

func (db *DB) updateHelper(collectionName, ID string, info bson.M) *mongo.SingleResult {
	collection, ctx := db.ctxDeferHelper(collectionName)
	_id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": info}

	results := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	return results
}

func (db *DB) CreateChatboard(input *model.NewChatboard) (*model.Chatboard, error) {
	res, err := db.resErrHelper("chatboards", input)

	chatboard := &model.Chatboard{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}

	return chatboard, err
}

func (db *DB) CreateMessage(input *model.NewMessage) (*model.Message, error) {
	res, err := db.resErrHelper("messages", input)

	message := &model.Message{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}

	return message, err
}

func (db *DB) GetMessages(ID string) ([]*model.Message, error) {
	res, ctx := db.multipleFetchHelper("messages", ID, "messageBy")
	var (
		messages []*model.Message
		err      error
	)

	defer res.Close(ctx)

	if err = res.All(context.TODO(), messages); err != nil {
		panic(err)
	}

	return messages, err
}

func (db *DB) SingleChatboard(ID string) (*model.Chatboard, error) {
	collection, ctx := db.ctxDeferHelper("chatboards")
	var chatboard *model.Chatboard

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&chatboard)

	return chatboard, err
}

func (db *DB) UpdateChatboard(ID string, input *model.UpdateChatboard) (*model.Chatboard, error) {
	var chatboard *model.Chatboard

	updateInfo := bson.M{}

	if input.Description != nil {
		updateInfo["description"] = input.Description
	}
	if input.ImageURL != nil {
		updateInfo["imageURL"] = input.ImageURL
	}
	if input.Name != nil {
		updateInfo["name"] = input.Name
	}

	results := db.updateHelper("chatboards", ID, updateInfo)
	if err := results.Decode(&chatboard); err != nil {
		log.Fatal(err)
		return chatboard, err
	}

	return chatboard, nil
}

func (db *DB) UpdateMessage(ID string, input *model.UpdateMessage) (*model.Message, error) {
	var message *model.Message

	updateInfo := bson.M{}

	if input.FileURL != nil {
		updateInfo["fileURL"] = input.FileURL
	}
	if input.Text != nil {
		updateInfo["text"] = input.Text
	}

	results := db.updateHelper("messages", ID, updateInfo)
	if err := results.Decode(&message); err != nil {
		log.Fatal(err)
		return message, err
	}

	return message, nil
}

func (db *DB) DeleteChatboard(ID string) (*model.DeleteChatboard, error) {
	err := db.deleteHelper("chatboards", ID)
	var delete *model.DeleteChatboard

	return delete, err
}

func (db *DB) DeleteMessage(ID string) (*model.DeleteMessage, error) {
	err := db.deleteHelper("messages", ID)
	var delete *model.DeleteMessage

	return delete, err
}

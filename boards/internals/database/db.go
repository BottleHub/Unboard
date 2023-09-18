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

func (db *DB) multipleFetchHelper(collectionName string) (*mongo.Cursor, context.Context) {
	collection, ctx := db.ctxDeferHelper(collectionName)

	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return res, ctx
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

func (db *DB) GetChatboards() ([]*model.Chatboard, error) {
	res, ctx := db.multipleFetchHelper("chatboards")
	var (
		chatboards []*model.Chatboard
		err        error
	)

	defer res.Close(ctx)

	if err = res.All(context.TODO(), &chatboards); err != nil {
		panic(err)
	}

	return chatboards, err
}

func (db *DB) GetMessages() ([]*model.Message, error) {
	res, ctx := db.multipleFetchHelper("messages")
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

func (db *DB) SingleMessage(ID string) (*model.Message, error) {
	collection, ctx := db.ctxDeferHelper("messages")
	var message *model.Message

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&message)

	return message, err
}

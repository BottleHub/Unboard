package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bottlehub/unboard/graph/model"

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

func ConnectDB() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))
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
	return &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("UserBase").Collection(collectionName)
}

func (db *DB) ctxDeferHelper(collectionName string) (*mongo.Collection, context.Context) {
	collection := colHelper(db, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

func (db *DB) multipleFetchHelper(collectionName string) *mongo.Cursor {
	collection, ctx := db.ctxDeferHelper(collectionName)

	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)

	return res
}

func (db *DB) CreateComment(input *model.NewComment) (*model.Comment, error) {
	res, err := db.resErrHelper("comments", input)

	comment := &model.Comment{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Text:      input.Text,
		CommentBy: &model.User{},
		CommentOn: &model.Post{},
	}

	return comment, err
}

func (db *DB) CreatePost(input *model.NewPost) (*model.Post, error) {
	res, err := db.resErrHelper("posts", input)

	post := &model.Post{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		PostedBy:    &model.User{},
		ImageURL:    input.ImageURL,
		Description: input.Description,
		Likes:       input.Likes,
	}

	return post, err
}

func (db *DB) CreateUser(input *model.NewUser) (*model.User, error) {
	password, _ := HashPassword(input.Password)
	input.Password = password

	res, err := db.resErrHelper("users", input)

	user := &model.User{
		ID:             res.InsertedID.(primitive.ObjectID).Hex(),
		Username:       input.Username,
		Name:           input.Name,
		About:          input.About,
		Email:          input.Email,
		AvatarImageURL: input.AvatarImageURL,
		Password:       password,
	}

	return user, err
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

func (db *DB) GetComments() ([]*model.Comment, error) {
	res := db.multipleFetchHelper("comments")
	var (
		comments []*model.Comment
		err      error
	)

	if err = res.All(context.TODO(), comments); err != nil {
		panic(err)
	}

	return comments, err
}

func (db *DB) GetPosts() ([]*model.Post, error) {
	res := db.multipleFetchHelper("posts")
	var (
		posts []*model.Post
		err   error
	)

	if err = res.All(context.TODO(), posts); err != nil {
		panic(err)
	}

	return posts, err
}

func (db *DB) GetUsers() ([]*model.User, error) {
	res := db.multipleFetchHelper("users")
	var (
		users []*model.User
		err   error
	)

	if err = res.All(context.TODO(), users); err != nil {
		panic(err)
	}

	return users, err
}

func (db *DB) GetChatboards() ([]*model.Chatboard, error) {
	res := db.multipleFetchHelper("chatboards")
	var (
		chatboards []*model.Chatboard
		err        error
	)

	if err = res.All(context.TODO(), chatboards); err != nil {
		panic(err)
	}

	return chatboards, err
}

func (db *DB) GetMessages() ([]*model.Message, error) {
	res := db.multipleFetchHelper("messages")
	var (
		messages []*model.Message
		err      error
	)

	if err = res.All(context.TODO(), messages); err != nil {
		panic(err)
	}

	return messages, err
}

func (db *DB) SingleComment(ID string) (*model.Comment, error) {
	collection, ctx := db.ctxDeferHelper("comments")
	var comment *model.Comment

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&comment)

	return comment, err
}

func (db *DB) SinglePost(ID string) (*model.Post, error) {
	collection, ctx := db.ctxDeferHelper("posts")
	var post *model.Post

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&post)

	return post, err
}

func (db *DB) SingleUser(ID string) (*model.User, error) {
	collection, ctx := db.ctxDeferHelper("users")
	var user *model.User

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	return user, err
}

func (db *DB) GetUserIdByUsername(username string) (string, error) {
	collection, ctx := db.ctxDeferHelper("users")
	var user *model.User

	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	id := user.ID
	return id, err
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

func (db *DB) DeletedComment(ID string) (*model.DeleteComment, error) {
	collection, ctx := db.ctxDeferHelper("comments")

	_id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": _id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteComment{ID: ID}, err
}
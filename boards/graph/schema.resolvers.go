package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"github.com/bottlehub/unboard/boards/graph/model"
	"github.com/bottlehub/unboard/boards/internal/database"
)

// CreateChatboard is the resolver for the createChatboard field.
func (r *mutationResolver) CreateChatboard(ctx context.Context, input model.NewChatboard) (*model.Chatboard, error) {
	chatboard, err := db.CreateChatboard(&input)
	return chatboard, err
}

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	message, err := db.CreateMessage(&input)
	return message, err
}

// UpdateChatboard is the resolver for the updateChatboard field.
func (r *mutationResolver) UpdateChatboard(ctx context.Context, id string, input model.UpdateChatboard) (*model.Chatboard, error) {
	chatboard, err := db.UpdateChatboard(id, &input)
	return chatboard, err
}

// UpdateMessage is the resolver for the updateMessage field.
func (r *mutationResolver) UpdateMessage(ctx context.Context, id string, input model.UpdateMessage) (*model.Message, error) {
	message, err := db.UpdateMessage(id, &input)
	return message, err
}

// DeleteChatboard is the resolver for the deleteChatboard field.
func (r *mutationResolver) DeleteChatboard(ctx context.Context, id string) (*model.DeleteChatboard, error) {
	chatboard, err := db.DeleteChatboard(id)
	return chatboard, err
}

// DeleteMessage is the resolver for the deleteMessage field.
func (r *mutationResolver) DeleteMessage(ctx context.Context, id string) (*model.DeleteMessage, error) {
	messages, err := db.DeleteMessage(id)
	return messages, err
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context, input model.Fetch) ([]*model.Message, error) {
	messages, err := db.GetMessages(input.ID)
	return messages, err
}

// Chatboard is the resolver for the chatboard field.
func (r *queryResolver) Chatboard(ctx context.Context, input model.Fetch) (*model.Chatboard, error) {
	chatboard, err := db.SingleChatboard(input.ID)
	return chatboard, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	db, _ = database.ConnectDB()
)

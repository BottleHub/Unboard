package graph

// import (
// 	"context"

// 	"github.com/bottlehub/unboard/backend/graph/model"
// )

// // CreateProject is the resolver for the createProject field.
// func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
// 	post, err := db.CreatePost(&input)
// 	return post, err
// }

// // CreateProject is the resolver for the createProject field.
// func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
// 	comment, err := db.CreateComment(&input)
// 	return comment, err
// }

// // CreateProject is the resolver for the createProject field.
// func (r *mutationResolver) CreateChatboard(ctx context.Context, input model.NewChatboard) (*model.Chatboard, error) {
// 	chatboard, err := db.CreateChatboard(&input)
// 	return chatboard, err
// }

// // CreateProject is the resolver for the createProject field.
// func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
// 	message, err := db.CreateMessage(&input)
// 	return message, err
// }

// // CreateLink is the resolver for the createLink field.
// func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
// 	var link model.Link
// 	var user model.User
// 	link.Address = input.Address
// 	link.Title = input.Title
// 	user.Name = "test"
// 	link.User = &user
// 	return &link, nil
// }

// // Login implements MutationResolver.
// func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
// 	panic("unimplemented")
// }

// // RefreshToken implements MutationResolver.
// func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
// 	panic("unimplemented")
// }

// // Owners is the resolver for the owners field.
// func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
// 	posts, err := db.GetPosts()
// 	return posts, err
// }

// // Owners is the resolver for the owners field.
// func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
// 	comments, err := db.GetComments()
// 	return comments, err
// }

// // Owners is the resolver for the owners field.
// func (r *queryResolver) Chatboards(ctx context.Context) ([]*model.Chatboard, error) {
// 	chatboard, err := db.GetChatboards()
// 	return chatboard, err
// }

// // Owners is the resolver for the owners field.
// func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
// 	messages, err := db.GetMessages()
// 	return messages, err
// }

// // Owner is the resolver for the owner field.
// func (r *queryResolver) Post(ctx context.Context, input *model.FetchPost) (*model.Post, error) {
// 	post, err := db.SinglePost(input.ID)
// 	return post, err
// }

// // Owner is the resolver for the owner field.
// func (r *queryResolver) Comment(ctx context.Context, input model.FetchComment) (*model.Comment, error) {
// 	comment, err := db.SingleComment(input.ID)
// 	return comment, err
// }

// // Owner is the resolver for the owner field.
// func (r *queryResolver) Chatboard(ctx context.Context, input model.FetchChatboard) (*model.Chatboard, error) {
// 	chatboard, err := db.SingleChatboard(input.ID)
// 	return chatboard, err
// }

// // Owner is the resolver for the owner field.
// func (r *queryResolver) Message(ctx context.Context, input model.FetchMessage) (*model.Message, error) {
// 	message, err := db.SingleMessage(input.ID)
// 	return message, err
// }

// // Links is the resolver for the links field.
// func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
// 	var links []*model.Link
// 	dummyLink := model.Link{
// 		Title:   "our dummy link",
// 		Address: "https://address.org",
// 		User:    &model.User{Name: "admin"},
// 	}
// 	links = append(links, &dummyLink)
// 	return links, nil
// }

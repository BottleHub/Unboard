// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Chatboard struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	ImageURL    string     `json:"imageURL"`
	Description *string    `json:"description,omitempty"`
	Members     []*User    `json:"members,omitempty"`
	Messages    []*Message `json:"messages,omitempty"`
}

type DeleteChatboard struct {
	ID string `json:"id"`
}

type DeleteMessage struct {
	ID string `json:"id"`
}

type Fetch struct {
	ID string `json:"id"`
}

type FetchChatboards struct {
	UserID string `json:"userId"`
}

type FetchMessages struct {
	ChatboardID string `json:"chatboardId"`
}

type Message struct {
	ID        string     `json:"id"`
	Text      *string    `json:"text,omitempty"`
	FileURL   *string    `json:"fileURL,omitempty"`
	MessageBy *User      `json:"messageBy"`
	MessageOn *Chatboard `json:"messageOn"`
}

type NewChatboard struct {
	Name        string  `json:"name"`
	ImageURL    string  `json:"imageURL"`
	Description *string `json:"description,omitempty"`
}

type NewMessage struct {
	Text      *string `json:"text,omitempty"`
	FileURL   *string `json:"fileURL,omitempty"`
	MessageBy string  `json:"messageBy"`
	MessageOn string  `json:"messageOn"`
}

type UpdateChatboard struct {
	Name        *string `json:"name,omitempty"`
	ImageURL    *string `json:"imageURL,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateMessage struct {
	Text    *string `json:"text,omitempty"`
	FileURL *string `json:"fileURL,omitempty"`
}

type User struct {
	ID             string  `json:"id"`
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	About          *string `json:"about,omitempty"`
	AvatarImageURL string  `json:"avatarImageURL"`
	Following      []*User `json:"following,omitempty"`
	Followers      []*User `json:"followers,omitempty"`
}

type Status string

const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

var AllStatus = []Status{
	StatusNotStarted,
	StatusInProgress,
	StatusCompleted,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusNotStarted, StatusInProgress, StatusCompleted:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

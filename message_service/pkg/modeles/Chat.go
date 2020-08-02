package modeles

import "time"

type Chat struct {
	ChatId    string       `json:"id"`
	Name      string       `json:"name"`
	Users     []User  	   `json:"users"`
	CreatedAt time.Time    `json:"created_at"`
}

type ChatAddRequest struct {
	Name      string  `json:"name"`
	UsersId   []int   `json:"users"`
}

type ChatAddResponse struct {
	ChatId      int  `json:"id"`
}

type ChatsGetRequest struct {
	UserId      string  `json:"user"`
}

type ChatsGetResponse struct {
	Chats      []Chat  `json:"chats"`
}

func (s ChatsGetResponse) Less(i, j int) bool {
	return s.Chats[i].CreatedAt.Before(s.Chats[j].CreatedAt)
}

func (s ChatsGetResponse) Swap(i, j int) {
	s.Chats[i].CreatedAt,
		s.Chats[j].CreatedAt = s.Chats[j].CreatedAt,
		s.Chats[i].CreatedAt
}

func (s ChatsGetResponse) Len() int {
	return len(s.Chats)
}
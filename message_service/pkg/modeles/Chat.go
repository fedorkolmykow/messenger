package modeles

import "time"

type Chat struct {
	ChatId    string       `json:"id"`
	Name      string       `json:"name"`
	Users     []string     `json:"users"`
	CreatedAt string       `json:"created_at"`
}

type ChatAddRequest struct {
	Name      string       `json:"name"`
	UsersId   []string     `json:"users"`
	CreatedAt string       `json:"-"`
}

type ChatAddResponse struct {
	ChatId      string  `json:"id"`
}

type ChatsGetRequest struct {
	UserId      string  `json:"user"`
}

type ChatsGetResponse struct {
	Chats      []Chat  `json:"chats"`
}

func (s ChatsGetResponse) Less(i, j int) bool {
	ti, _ := time.Parse(time.RFC822 ,s.Chats[i].CreatedAt)
	tj, _ := time.Parse(time.RFC822 ,s.Chats[j].CreatedAt)
	return ti.Before(tj)
}

func (s ChatsGetResponse) Swap(i, j int) {
	s.Chats[i].CreatedAt, s.Chats[j].CreatedAt =
		s.Chats[j].CreatedAt, s.Chats[i].CreatedAt
}

func (s ChatsGetResponse) Len() int {
	return len(s.Chats)
}
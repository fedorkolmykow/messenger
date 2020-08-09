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

func (cgr ChatsGetResponse) Less(i, j int) bool {
	ti, _ := time.Parse(time.RFC822 ,cgr.Chats[i].CreatedAt)
	tj, _ := time.Parse(time.RFC822 ,cgr.Chats[j].CreatedAt)
	return ti.Before(tj)
}

func (cgr ChatsGetResponse) Swap(i, j int) {
	cgr.Chats[i].CreatedAt, cgr.Chats[j].CreatedAt =
		cgr.Chats[j].CreatedAt, cgr.Chats[i].CreatedAt
}

func (cgr ChatsGetResponse) Len() int {
	return len(cgr.Chats)
}
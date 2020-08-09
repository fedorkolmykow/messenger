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

type Chats []Chat

type ChatsGetResponse struct {
	Chats      Chats  `json:"chats"`
}

func (c Chats) Less(i, j int) bool {
	ti, _ := time.Parse(time.RFC822 ,c[i].CreatedAt)
	tj, _ := time.Parse(time.RFC822 ,c[j].CreatedAt)
	return ti.Before(tj)
}

func (c Chats) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Chats) Len() int {
	return len(c)
}
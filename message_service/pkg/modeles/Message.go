package modeles

import "time"

type Message struct {
	MessageId    string     `json:"id"`
	ChatId       string     `json:"chat"`
	AuthorId     string     `json:"author"`
	Text         string     `json:"text"`
	CreatedAt    time.Time  `json:"created_at"`
}

type MessageAddRequest struct {
	ChatId      string  `json:"chat"`
	AuthorId    string  `json:"author"`
	Text        string  `json:"text"`
}

type MessageAddResponse struct {
	MessageId      int  `json:"id"`
}

type MessagesGetRequest struct {
	ChatId      string  `json:"chat"`
}

type MessagesGetResponse struct {
	Messages      []Message  `json:"messages"`
}

func (s MessagesGetResponse) Less(i, j int) bool {
	return s.Messages[i].CreatedAt.Before(s.Messages[j].CreatedAt)
}

func (s MessagesGetResponse) Swap(i, j int) {
	s.Messages[i].CreatedAt,
	s.Messages[j].CreatedAt = s.Messages[j].CreatedAt,
	s.Messages[i].CreatedAt
}

func (s MessagesGetResponse) Len() int {
	return len(s.Messages)
}
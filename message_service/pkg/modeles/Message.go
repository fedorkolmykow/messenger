package modeles

import "time"

type Message struct {
	MessageId    string     `json:"id"`
	ChatId       string     `json:"chat"`
	AuthorId     string     `json:"author"`
	Text         string     `json:"text"`
	CreatedAt    string     `json:"created_at"`
}

type MessageAddRequest struct {
	ChatId      string     `json:"chat"`
	AuthorId    string     `json:"author"`
	Text        string     `json:"text"`
	CreatedAt   string     `json:"-"`
}

type MessageAddResponse struct {
	MessageId      string  `json:"id"`
}

type MessagesGetRequest struct {
	ChatId      string     `json:"chat"`
}

type MessagesGetResponse struct {
	Messages    []Message  `json:"messages"`
}

func (s MessagesGetResponse) Less(i, j int) bool {
	ti, _ := time.Parse(time.RFC822 ,s.Messages[i].CreatedAt)
	tj, _ := time.Parse(time.RFC822 ,s.Messages[j].CreatedAt)
	return ti.Before(tj)
}

func (s MessagesGetResponse) Swap(i, j int) {
	s.Messages[i].CreatedAt, s.Messages[j].CreatedAt =
		s.Messages[j].CreatedAt, s.Messages[i].CreatedAt
}

func (s MessagesGetResponse) Len() int {
	return len(s.Messages)
}
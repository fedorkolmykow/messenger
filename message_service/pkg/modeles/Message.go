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

func (mgr MessagesGetResponse) Less(i, j int) bool {
	ti, _ := time.Parse(time.RFC822 ,mgr.Messages[i].CreatedAt)
	tj, _ := time.Parse(time.RFC822 ,mgr.Messages[j].CreatedAt)
	return ti.Before(tj)
}

func (mgr MessagesGetResponse) Swap(i, j int) {
	mgr.Messages[i].CreatedAt, mgr.Messages[j].CreatedAt =
		mgr.Messages[j].CreatedAt, mgr.Messages[i].CreatedAt
}

func (mgr MessagesGetResponse) Len() int {
	return len(mgr.Messages)
}
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

type Messages []Message

type MessagesGetResponse struct {
	Messages    Messages  `json:"messages"`
}

func (m Messages) Less(i, j int) bool {
	ti, _ := time.Parse(time.RFC822 ,m[i].CreatedAt)
	tj, _ := time.Parse(time.RFC822 ,m[j].CreatedAt)
	return ti.Before(tj)
}

func (m Messages) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Messages) Len() int {
	return len(m)
}
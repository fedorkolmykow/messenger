package modeles

type Message struct {
	MessageId    string  `json:"id"`
	ChatId       string  `json:"chat"`
	AuthorId     string  `json:"author"`
	Text         string  `json:"text"`
	CreatedAt    string  `json:"created_at"`
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

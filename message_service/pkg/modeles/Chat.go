package modeles

type Chat struct {
	ChatId    string  `json:"id"`
	Name      string  `json:"name"`
	Users     []User  `json:"users"`
	CreatedAt string  `json:"created_at"`
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


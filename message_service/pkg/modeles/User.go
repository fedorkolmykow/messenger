package modeles

type User struct {
	UserId     string   `json:"id"`
	Username   string   `json:"username"`
	CreatedAt  string   `json:"created_at"`
}

type UserAddRequest struct {
	Username   string       `json:"username"`
	CreatedAt  string       `json:"-"`
}

type UserAddResponse struct {
	UserId     string      `json:"id"`
}
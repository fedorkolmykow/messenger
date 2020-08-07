package httpserver

import (
	m "avito_message/message_service/pkg/modeles"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

const(
	userAdd = iota
	chatAdd
	messageAdd
	chatsGet
	messagesGet
)

type correctService struct{
}

type errorService struct{
}

type TestCase struct {
	Req     []byte
	Resp    string
	Status  int
	S       server
	Handle  int
}

func TestHandle(t *testing.T){
	cases := []TestCase{
		{
			Req:          []byte(`{"username":"Alice"}`),
			Resp:         `{"id":"0"}`,
			Status:       http.StatusOK,
			S:            server{svc: &correctService{}},
			Handle:       userAdd,
		},
		{
			Req:          []byte(`username:Alice`),
			Resp:         ``,
			Status:       http.StatusBadRequest,
			S:            server{svc: &correctService{}},
			Handle:       userAdd,
		},
		{
			Req:          []byte(`{}`),
			Resp:         ``,
			Status:       http.StatusInternalServerError,
			S:            server{svc: &errorService{}},
			Handle:       userAdd,
		},
		{
			Req:          []byte(`{"name":"newChat", "users": ["0", "1", "2"]}`),
			Resp:         `{"id":"0"}`,
			Status:       http.StatusOK,
			S:            server{svc: &correctService{}},
			Handle:       chatAdd,
		},
		{
			Req:          []byte(`"name":"newChat", "users": ["0", "1", "2"]`),
			Resp:         ``,
			Status:       http.StatusBadRequest,
			S:            server{svc: &correctService{}},
			Handle:       chatAdd,
		},
		{
			Req:          []byte(`{}`),
			Resp:         ``,
			Status:       http.StatusInternalServerError,
			S:            server{svc: &errorService{}},
			Handle:       chatAdd,
		},
		{
			Req:          []byte(`{"chat": "1", "author": "0", "text": "HELLO!"}`),
			Resp:         `{"id":"0"}`,
			Status:       http.StatusOK,
			S:            server{svc: &correctService{}},
			Handle:       messageAdd,
		},
		{
			Req:          []byte(`"chat": "1", "author": "0", "text": "HELLO!"`),
			Resp:         ``,
			Status:       http.StatusBadRequest,
			S:            server{svc: &correctService{}},
			Handle:       messageAdd,
		},
		{
			Req:          []byte(`{}`),
			Resp:         ``,
			Status:       http.StatusInternalServerError,
			S:            server{svc: &errorService{}},
			Handle:       messageAdd,
		},
		//{
		//	Req:          []byte(`{"user":"0"}`),
		//	Resp:         `{"chats": [{"id": "0", "name":"newChat", "users": ["0", "1", "2"], "created_at": "2012.02.01"}]}`,
		//	Status:       http.StatusInternalServerError,
		//	S:            server{svc: &errorService{}},
		//	Handle:       chatsGet,
		//},
	}
	for num, c := range cases{
		req := httptest.NewRequest("POST", "http://localhost/user", bytes.NewBuffer(c.Req))
		w := httptest.NewRecorder()
		switch c.Handle {
		    case userAdd:     c.S.HandleAddUser(w, req)
		    case chatAdd:     c.S.HandleAddChat(w, req)
		    case messageAdd:  c.S.HandleAddMessage(w, req)
		    case chatsGet:    c.S.HandleGetChat(w, req)
		    case messagesGet: c.S.HandleGetMessage(w, req)
		}

		if w.Result().StatusCode != c.Status{
			t.Errorf("[%d] unexpected status: %d, expected: %d",num, w.Result().StatusCode,  c.Status)
		}
		if c.Status == http.StatusOK{
			if c.Resp != w.Body.String(){
				t.Errorf("[%d] unexpected result:\n%s\nexpected:\n%s ", num, w.Body.String(), c.Resp)
			}
		}
	}
}

//correctService
func (s *correctService) AddUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error){
	return &m.UserAddResponse{UserId:"0"}, nil
}


func (s *correctService) AddChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error){
	return &m.ChatAddResponse{ChatId:"0"}, nil
}


func (s *correctService) AddMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error){
	return &m.MessageAddResponse{MessageId:"0"}, nil
}


func (s *correctService) GetChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error){
	return &m.ChatsGetResponse{}, nil
}


func (s *correctService) GetMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error){
	return &m.MessagesGetResponse{}, nil
}

//errorService
func (s *errorService) AddUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error){
	return nil, errors.New("Test error")
}

func (s *errorService) AddChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error){
	return nil, errors.New("Test error")
}

func (s *errorService) AddMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error){
	return nil, errors.New("Test error")
}

func (s *errorService) GetChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error){
	return nil, errors.New("Test error")
}

func (s *errorService) GetMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error){
	return nil, errors.New("Test error")
}

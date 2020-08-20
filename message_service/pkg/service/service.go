package service

import (
	m "github.com/fedorkolmykow/messesnger/pkg/modeles"
	"sort"
	"time"
)

type dbClient interface {
	InsertUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error)
	InsertChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error)
	InsertMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error)
	SelectChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error)
	SelectMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error)
}

type Service interface {
	AddUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error)
	AddChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error)
	AddMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error)
	GetChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error)
	GetMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error)
}

type service struct{
	db dbClient
}

func (s *service) AddUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error){
	userAddReq.CreatedAt = time.Now().Format(time.RFC822)
	return s.db.InsertUser(userAddReq)
}


func (s *service) AddChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error){
	chatAddReq.CreatedAt = time.Now().Format(time.RFC822)
	return s.db.InsertChat(chatAddReq)
}


func (s *service) AddMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error){
	mesAddReq.CreatedAt = time.Now().Format(time.RFC822)
	return s.db.InsertMessage(mesAddReq)
}


func (s *service) GetChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error){
	chatsGetResp, err = s.db.SelectChats(chatsGetReq)
	if chatsGetResp!=nil {
		sort.Sort(chatsGetResp.Chats)
	}
	return 
}


func (s *service) GetMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error){
	mesGetResp, err = s.db.SelectMessages(mesGetReq)
	if mesGetResp != nil{
		sort.Sort(mesGetResp.Messages)
	}
	return
}

func NewService(db dbClient) Service{
	return &service{db: db}
}
package service

import (
	m "avito_message/message_service/pkg/modeles"
	"sort"
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
	db dbclient
}

func (s *service) AddUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error){
	return s.db.InsertUser(userAddReq)
}


func (s *service) AddChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error){
	return s.db.InsertChat(chatAddReq)
}


func (s *service) AddMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error){
	return s.db.InsertMessage(mesAddReq)
}


func (s *service) GetChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error){
	chatsGetResp, err = s.db.SelectChats(chatsGetReq)
	sort.Sort(chatsGetResp)
	return 
}


func (s *service) GetMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error){
	mesGetResp, err = s.db.SelectMessages(mesGetReq)
	sort.Sort(mesGetResp)
	return
}

func NewService(db dbclient) Service{
	return &service{db: db}
}
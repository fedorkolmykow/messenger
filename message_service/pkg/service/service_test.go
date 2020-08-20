package service

import (
	"reflect"
	"testing"

	m "github.com/fedorkolmykow/messesnger/pkg/modeles"
)

type mockDbCon struct{
}

func Test_SortingChatsCorrect(t *testing.T){
	exp:=  []m.Chat{
		{
			ChatId:    "0",
			Name:      "Rats",
			Users:     []string{"0", "2"},
			CreatedAt: "08 Aug 20 16:49 UTC",
		},
		{
			ChatId:    "2",
			Name:      "Teachers",
			Users:     []string{"0", "3"},
			CreatedAt: "09 Aug 20 16:49 UTC",
		},
		{
			ChatId:    "3",
			Name:      "World government",
			Users:     []string{"0","2","3"},
			CreatedAt: "10 Aug 20 16:49 UTC",
		},
	}
	svc := NewService(&mockDbCon{})
    res, err:= svc.GetChats(&m.ChatsGetRequest{})
    if err != nil{
		t.Errorf("Unexpected error: %s", err)
	}

	for i := range exp{
		if !reflect.DeepEqual(exp[i],res.Chats[i]){
			t.Errorf("\nUnexpected result:\n%s\nExpected:\n%s ", res.Chats, exp)
			break
		}
	}
	//if !reflect.DeepEqual(exp,res.Chats){
	//	t.Errorf("\nUnexpected result:\n%s\nExpected:\n%s ", res.Chats, exp)
	//}
}

func Test_SortingMessagesCorrect(t *testing.T){
	exp:=  []m.Message{
		{
			MessageId: "1",
			ChatId:    "3",
			AuthorId:  "0",
			Text:      "Is it a bird?",
			CreatedAt: "09 Aug 20 13:49 UTC",
		},
		{
			MessageId: "2",
			ChatId:    "3",
			AuthorId:  "2",
			Text:      "Is it a plane?",
			CreatedAt: "09 Aug 20 13:50 UTC",
		},
		{
			MessageId: "3",
			ChatId:    "3",
			AuthorId:  "3",
			Text:      "No, it is Superman!",
			CreatedAt: "09 Aug 20 13:51 UTC",
		},
	}
	svc := NewService(&mockDbCon{})
	res, err:= svc.GetMessages(&m.MessagesGetRequest{})
	if err != nil{
		t.Errorf("Unexpected error: %s", err)
	}

	for i := range exp{
		if !reflect.DeepEqual(exp[i],res.Messages[i]){
			t.Errorf("\nUnexpected result:\n%s\nExpected:\n%s ", res.Messages, exp)
			break
		}
	}
}



func (mock *mockDbCon) SelectChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error){
	return &m.ChatsGetResponse{
		Chats: []m.Chat{
			{
				ChatId:    "3",
				Name:      "World government",
				Users:     []string{"0","2","3"},
				CreatedAt: "10 Aug 20 16:49 UTC",
			},
			{
				ChatId:    "0",
				Name:      "Rats",
				Users:     []string{"0", "2"},
				CreatedAt: "08 Aug 20 16:49 UTC",
			},
			{
				ChatId:    "2",
				Name:      "Teachers",
				Users:     []string{"0", "3"},
				CreatedAt: "09 Aug 20 16:49 UTC",
			},
		},
	}, nil// errors.New("Test error")
}
func (mock *mockDbCon) SelectMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error){
	return &m.MessagesGetResponse{
		Messages: []m.Message{
			{
				MessageId: "3",
				ChatId:    "3",
				AuthorId:  "3",
				Text:      "No, it is Superman!",
				CreatedAt: "09 Aug 20 13:51 UTC",
			},
			{
				MessageId: "1",
				ChatId:    "3",
				AuthorId:  "0",
				Text:      "Is it a bird?",
				CreatedAt: "09 Aug 20 13:49 UTC",
			},
			{
				MessageId: "2",
				ChatId:    "3",
				AuthorId:  "2",
				Text:      "Is it a plane?",
				CreatedAt: "09 Aug 20 13:50 UTC",
			},
		},
	}, nil
}

func (mock *mockDbCon) InsertUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error){return}
func (mock *mockDbCon) InsertChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error){return}
func (mock *mockDbCon) InsertMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error){return}
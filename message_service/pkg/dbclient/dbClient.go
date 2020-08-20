package dbclient

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4"

	m "github.com/fedorkolmykow/messesnger/pkg/modeles"
)

type DbClient interface {
	InsertUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error)
	InsertChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error)
	InsertMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error)
	SelectChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error)
	SelectMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error)
}

type db struct {
	dbCon *pgx.Conn
}

const(
insertUser = "INSERT INTO Users (username, created_at) VALUES ($1, $2) RETURNING user_id;"
insertChat = "INSERT INTO Chats (name, created_at) VALUES ($1, $2) RETURNING chat_id;"
insertChatUsers = "INSERT INTO Chat_Users (chat_id, user_id) VALUES ($1, $2);"
insertMessage = "INSERT INTO Messages (author_id, chat_id, text, created_at) " +
	"VALUES ($1, $2, $3, $4) RETURNING message_id;"
selectChats = "SELECT Chats.chat_id, Chats.name, Chats.created_at," +
	" ARRAY(SELECT user_id FROM Chat_Users WHERE chat_id=Chats.chat_id) as Users" +
	" FROM Chats " +
	" LEFT JOIN Chat_Users ON Chat_Users.chat_id = Chats.chat_id" +
	" WHERE Chat_Users.user_id = $1;"
selectMessages = "SELECT message_id, author_id, chat_id, text, created_at " +
	"FROM Messages " +
	"WHERE chat_id = $1;"
)

func (d *db) InsertUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error) {
	var userId int
	userAddResp = &m.UserAddResponse{}
	row := d.dbCon.QueryRow(context.Background(), insertUser, userAddReq.Username, userAddReq.CreatedAt)
	err = row.Scan(&userId)
	if err!=nil{
		log.Println(err)
		return
	}
	userAddResp.UserId = strconv.Itoa(userId)
	return
}
func (d *db) InsertChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error) {
	var chatId int
	b := &pgx.Batch{}
	chatAddResp = &m.ChatAddResponse{}
	row := d.dbCon.QueryRow(context.Background(), insertChat, chatAddReq.Name, chatAddReq.CreatedAt)
	err = row.Scan(&chatId)
	if err!=nil{
		log.Println(err)
		return
	}
	chatAddResp.ChatId = strconv.Itoa(chatId)
	for i, _ := range chatAddReq.UsersId{
		b.Queue(insertChatUsers, chatAddResp.ChatId, chatAddReq.UsersId[i])
	}
	br := d.dbCon.SendBatch(context.Background(), b)
	defer br.Close()
	rows, err := br.Query()
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	return
}
func (d *db) InsertMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error) {
	var mesId int
	mesAddResp = &m.MessageAddResponse{}
	row := d.dbCon.QueryRow(context.Background(), insertMessage,
		mesAddReq.AuthorId, mesAddReq.ChatId, mesAddReq.Text, mesAddReq.CreatedAt)
	err = row.Scan(&mesId)
	if err!=nil{
		log.Println(err)
		return
	}
	mesAddResp.MessageId = strconv.Itoa(mesId)
	return
}
func (d *db) SelectChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error) {
	var chatId int
	var createdAt time.Time
	rows, err := d.dbCon.Query(context.Background(), selectChats, chatsGetReq.UserId)
	if err != nil{
		log.Println(err)
		return
	}
	defer rows.Close()
	chatsGetResp = &m.ChatsGetResponse{Chats: []m.Chat{}}
	for rows.Next(){
		c := m.Chat{}
		usersId := []int{}
		err = rows.Scan(&chatId, &c.Name, &createdAt, &usersId)
		if err != nil{
			log.Println(err)
			return
		}
		c.ChatId = strconv.Itoa(chatId)
		c.CreatedAt = createdAt.Format(time.RFC822)
		for i := range usersId {
			id := strconv.Itoa(usersId[i])
			c.Users = append(c.Users, id)
		}
		chatsGetResp.Chats = append(chatsGetResp.Chats, c)
	}

	return
}
func (d *db) SelectMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error) {
	var mesId, auId, chatId int
	var createdAt time.Time
	rows, err := d.dbCon.Query(context.Background(), selectMessages, mesGetReq.ChatId)
	if err != nil{
		log.Println(err)
		return
	}
	defer rows.Close()
	mesGetResp = &m.MessagesGetResponse{Messages: []m.Message{}}
	for rows.Next(){
		mes := m.Message{}
		err = rows.Scan(&mesId, &auId, &chatId, &mes.Text, &createdAt)
		if err != nil{
			log.Println(err)
			return
		}
		mes.MessageId = strconv.Itoa(mesId)
		mes.ChatId = strconv.Itoa(chatId)
		mes.AuthorId = strconv.Itoa(auId)
		mes.CreatedAt = createdAt.Format(time.RFC822)
		mesGetResp.Messages = append(mesGetResp.Messages, mes)
	}
	return
}

// NewDb returns a new Db instance.
func NewDb() DbClient {
	dbCon, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}

	return &db{dbCon: dbCon}
}
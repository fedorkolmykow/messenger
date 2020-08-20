package httpserver

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"

	m "github.com/fedorkolmykow/messesnger/pkg/modeles"
)

type service interface {
	AddUser(userAddReq *m.UserAddRequest) (userAddResp *m.UserAddResponse, err error)
	AddChat(chatAddReq *m.ChatAddRequest) (chatAddResp *m.ChatAddResponse, err error)
	AddMessage(mesAddReq *m.MessageAddRequest) (mesAddResp *m.MessageAddResponse, err error)
	GetChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error)
	GetMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error)
}

type server struct {
	svc service
}

func (s *server) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	userAddReq:= &m.UserAddRequest{}
	err = userAddReq.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userAddResp, err := s.svc.AddUser(userAddReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := userAddResp.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error() ,http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}

func (s *server) HandleAddChat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	chatAddReq:= &m.ChatAddRequest{}
	err = chatAddReq.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chatAddResp, err := s.svc.AddChat(chatAddReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := chatAddResp.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}

func (s *server) HandleGetChat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	chatsGetReq:= &m.ChatsGetRequest{}
	err = chatsGetReq.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chatsGetResp, err := s.svc.GetChats(chatsGetReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := chatsGetResp.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}

func (s *server) HandleAddMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	mesAddReq:= &m.MessageAddRequest{}
	err = mesAddReq.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mesAddResp, err := s.svc.AddMessage(mesAddReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := mesAddResp.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}

func (s *server) HandleGetMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	mesGetReq:= &m.MessagesGetRequest{}
	err = mesGetReq.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mesGetResp, err := s.svc.GetMessages(mesGetReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := mesGetResp.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Println(err)
	}
}

//func (s *server) Handle(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Привет, мир!")
//}

// NewServer returns a new mux.Router instance.
func NewServer(svc service) (httpServer *mux.Router) {
	s := server{svc: svc}
	router := mux.NewRouter()

	router.HandleFunc("/users/add", s.HandleAddUser).
		Methods("POST")

	router.HandleFunc("/chats/add", s.HandleAddChat).
		Methods("POST")

	router.HandleFunc("/chats/get", s.HandleGetChat).
		Methods("POST")

	router.HandleFunc("/messages/add", s.HandleAddMessage).
		Methods("POST")

	router.HandleFunc("/messages/get", s.HandleGetMessage).
		Methods("POST")

	return router
}

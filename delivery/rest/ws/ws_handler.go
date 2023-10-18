package ws

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/delivery/rest/helper"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/internal/_usecase"
)

type Handler struct {
	hub *_usecase.Hub
}

func NewHandler(h *_usecase.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	req := new(_usecase.RequestCreateRoom)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		helper.ErrorEncode(w, err)
		return
	}

	h.hub.Rooms[req.ID] = &_usecase.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*_usecase.Client),
	}

	helper.SuccessEncode(w, req, "created room successfully")
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := _usecase.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		helper.ErrorEncode(w, err)
		return
	}

	roomID := chi.URLParam(r, "room-id")
	clientID := r.URL.Query().Get("user-id")
	username := r.URL.Query().Get("username")

	cl := &_usecase.Client{
		Conn:     conn,
		Message:  make(chan *_usecase.Message, 10),
		ID:       clientID,
		Username: username,
		RoomID:   roomID,
	}

	m := &_usecase.Message{
		Content:  "A new user joined this room",
		RoomID:   roomID,
		Username: cl.Username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

func (h *Handler) GetRoom(w http.ResponseWriter, r *http.Request) {
	var rooms []_usecase.RoomResponse

	for _, room := range h.hub.Rooms {
		rooms = append(rooms, _usecase.RoomResponse{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	helper.SuccessEncode(w, rooms, "list rooms")
}

func (h *Handler) GetClient(w http.ResponseWriter, r *http.Request) {
	var clients []_usecase.ClientResponse
	roomID := chi.URLParam(r, "room-id")

	if _, ok := h.hub.Rooms[roomID]; !ok {
		helper.SuccessEncode(w, []_usecase.Client{}, "list rooms")
	}

	for _, client := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, _usecase.ClientResponse{
			ID:       client.ID,
			RoomID:   client.RoomID,
			Username: client.Username,
		})
	}

	helper.SuccessEncode(w, clients, "list clients")
}

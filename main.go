package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/delivery/rest"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/delivery/rest/ws"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/infra"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/internal/_repository"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/internal/_usecase"
)

func main() {
	db := infra.NewPostgresConnection()

	userRepo := _repository.NewUserRepository(db)

	userUsecase := _usecase.NewUserUsacaseImpl(userRepo)
	hub := _usecase.NewHub()

	go hub.Run()
	userHandler := rest.NewUserHandlerImpl(userUsecase)
	wsHandler := ws.NewHandler(hub)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/register", userHandler.Create)
	r.Post("/login", userHandler.Login)
	r.Post("/ws/create-room", wsHandler.CreateRoom)
	r.Get("/ws/join-room/{room-id}", wsHandler.JoinRoom)
	r.Get("/ws/get-room", wsHandler.GetRoom)
	r.Get("/ws/get-clients/{room-id}", wsHandler.GetClient)

	err := http.ListenAndServe(":8181", r)
	if err != nil {
		log.Fatalf("failed start http serve | err %v", err)
		return
	}
}

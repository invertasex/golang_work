package main

import (
	"encoding/json"
	"log"
	"net/http"

	"context"
	//"golang_work/grpc-service"

	"github.com/your-repo/grpc-service/server"

	// gRPC сгенерированный клиент

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type CheckUserRequest struct {
	BotToken    string `json:"bot_token"`
	ChannelLink string `json:"channel_link"`
	UserID      int64  `json:"user_id"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/check-user", checkUserHandler).Methods("POST")

	log.Println("HTTP-сервис запущен на 8080")
	http.ListenAndServe(":8080", r)
}

func checkUserHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Парсим входящие данные JSON
	var reqData CheckUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	// 2. Устанавливаем gRPC-соединение со вторым микросервисом
	conn, err := grpc.Dial("grpc-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Ошибка соединения с gRPC: %v", err)
		http.Error(w, "Не удалось связаться со gRPC-сервисом", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := server.NewTelegramServiceClient(conn)

	// 3. Вызываем RPC-метод
	grpcReq := &server.CheckUserRequest{
		BotToken:    reqData.BotToken,
		ChannelLink: reqData.ChannelLink,
		UserId:      reqData.UserID,
	}

	resp, err := client.CheckUser(context.Background(), grpcReq)
	if err != nil {
		log.Printf("Ошибка при вызове CheckUser: %v", err)
		http.Error(w, "Ошибка при обработке запроса во втором микросервисе", http.StatusInternalServerError)
		return
	}

	// 4. Возвращаем ответ клиенту
	response := map[string]interface{}{
		"is_in_group": resp.IsInGroup,
		"message":     resp.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

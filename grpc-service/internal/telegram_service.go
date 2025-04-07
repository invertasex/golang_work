package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"grpc_s/github.com/your-repo/grpc-service/server" // сгенерированные stubs
)

type telegramService struct {
	repo UserRepository
	server.UnimplementedTelegramServiceServer
}

func NewTelegramService(repo UserRepository) *telegramService {
	return &telegramService{
		repo: repo,
	}
}

// CheckUser — реализация RPC-метода
func (s *telegramService) CheckUser(ctx context.Context, req *server.CheckUserRequest) (*server.CheckUserResponse, error) {
	// 1. Проверим через Telegram Bot API, находится ли пользователь в группе:
	//    Пример использования: GET https://api.telegram.org/bot<BOT_TOKEN>/getChatMember?chat_id=<CHANNEL_ID>&user_id=<USER_ID>
	//    channel_link может быть вида "@testchannel" или "t.me/testchannel" и т.п.
	//    Нужно конвертировать channel_link -> chat_id (чаще всего, если это публичный канал, достаточно подставить как "@channelusername")
	//    Или напрямую передавать numeric ID канала, если известно.

	// Для примера предположим, что channel_link — это что-то вроде "@mychannel"
	// Можно убрать "@" при необходимости, но это детали.
	chatID := req.ChannelLink // упрощённо берем как есть

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getChatMember?chat_id=%s&user_id=%d",
		req.BotToken, chatID, req.UserId)

	resp, err := http.Get(url)
	if err != nil {
		return &server.CheckUserResponse{
			IsInGroup: false,
			Message:   "Ошибка запроса к Telegram API",
		}, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &server.CheckUserResponse{
			IsInGroup: false,
			Message:   "Ошибка чтения ответа от Telegram API",
		}, nil
	}

	// Парсим ответ. Пример JSON-ответа:
	// {
	//   "ok": true,
	//   "result": {
	//     "user": { "id": 123456, "is_bot": false, ... },
	//     "status": "member"
	//   }
	// }
	// Если "status" = "member", "administrator", "creator" и т.д. — значит пользователь в группе
	var tgResp map[string]interface{}
	err = json.Unmarshal(body, &tgResp)
	if err != nil {
		return &server.CheckUserResponse{
			IsInGroup: false,
			Message:   "Ошибка парсинга JSON от Telegram API",
		}, nil
	}

	// Примитивная проверка "status"
	isInGroup := false
	if result, ok := tgResp["result"].(map[string]interface{}); ok {
		if status, ok := result["status"].(string); ok {
			switch status {
			case "creator", "administrator", "member":
				isInGroup = true
			}
		}
	}

	if isInGroup {
		// Сохраняем пользователя в БД
		if err := s.repo.InsertUser(ctx, req.UserId, 0 /* channel_id; если нужен numeric, дополните */); err != nil {
			// Ошибка при вставке
			return &server.CheckUserResponse{
				IsInGroup: true,
				Message:   fmt.Sprintf("Пользователь в группе, но произошла ошибка вставки в БД: %v", err),
			}, nil
		}
		return &server.CheckUserResponse{
			IsInGroup: true,
			Message:   "Пользователь находится в группе и успешно добавлен в БД",
		}, nil
	}

	// Если пользователь не в группе
	return &server.CheckUserResponse{
		IsInGroup: false,
		Message:   "Пользователь не находится в группе",
	}, nil
}

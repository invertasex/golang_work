package main

import (
	"log"
	"net"

	"grpc_s/github.com/your-repo/grpc-service/server"
	"grpc_s/internal"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// Путь к сгенерированным protobuf файлам
)

func main() {
	// Подключаемся к базе
	dbPool, err := pgxpool.New(nil, "postgres://user:password@db:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer dbPool.Close()

	// Создаём репозиторий и сервис
	userRepo := internal.NewUserRepository(dbPool)
	telegramSvc := internal.NewTelegramService(userRepo)

	// Настраиваем gRPC-сервер
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при попытке начать слушать порт: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Регистрируем сервис
	server.RegisterTelegramServiceServer(grpcServer, telegramSvc)

	// Для отладки можно включить reflection
	reflection.Register(grpcServer)

	log.Println("gRPC-сервис запущен на порту 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Не удалось запустить gRPC-сервер: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "mock-cbr/docs"
	"mock-cbr/internal/handler"
	"mock-cbr/internal/proto"
	"mock-cbr/internal/storage"

	"google.golang.org/grpc"
)

// @title        Mock CBR API
// @version      1.0
// @description  Мок-сервис, имитирующий ответы ЦБ РФ
// @host         localhost:8080
// @BasePath     /
func main() {
	storage.InitDB()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Не удалось запустить gRPC: %v", err)
		}

		grpcServer := grpc.NewServer()
		proto.RegisterCurrencyServiceServer(grpcServer, &handler.CurrencyServiceServer{})

		fmt.Println("gRPC сервер запущен на порту :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Ошибка запуска gRPC: %v", err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/daily", handler.GetDailyHandler).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("HTTP сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
	}
}

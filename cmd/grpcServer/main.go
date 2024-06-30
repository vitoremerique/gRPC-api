package main

import (
	"database/sql"

	"net"
	"vitoremerique/gRPC/internal/database"
	"vitoremerique/gRPC/internal/pb"
	"vitoremerique/gRPC/internal/service"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, error := sql.Open("sqlite3", "./db.sqlite")
	if error != nil {
		panic(error)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	CategoryService := service.NewCategoryService(*categoryDB)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, CategoryService)

	lis, error := net.Listen("tcp", ":50051")
	if error != nil {
		panic(error)
	}

	if error := grpcServer.Serve(lis); error != nil {
		panic(error)
	}
}

package main

import (
	"database/sql"
	"github.com/Jhon-Henkel/go-lang-full-cycle-grpc/internal/database"
	"github.com/Jhon-Henkel/go-lang-full-cycle-grpc/internal/pb"
	"github.com/Jhon-Henkel/go-lang-full-cycle-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	CategoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, CategoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

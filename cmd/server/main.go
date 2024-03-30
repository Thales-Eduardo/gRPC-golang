package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/Thales-Eduardo/gRPC-golang/internal/database"
	"github.com/Thales-Eduardo/gRPC-golang/internal/pb"
	"github.com/Thales-Eduardo/gRPC-golang/internal/services"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// postgresql://postgres:docker@postgres_bd:5432/e-commerce?schema=public
	connStr := "postgres://postgres:docker@localhost:5432/graphql?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("falha ao conectar ao database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("falha ao pingar o database: %v", err)
	}
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso")

	categoryDB := database.NewCategory(db)
	// temos que registra o serviço no register do drpcServer ./internal/pb/*_grpc.pb.go
	categoryServices := services.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryServices)

	//reflesh server
	reflection.Register(grpcServer)

	//abrir uma porta tcp
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

package server

import (
	"fmt"
	"github.com/alpha-omega-corp/services/config"
	"github.com/alpha-omega-corp/services/database"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewGRPC(host string, env string, proto func(h *database.Handler, grpc *grpc.Server)) error {
	c := config.Get(env)
	listen, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	dbHandler := database.NewHandler(c.DB)

	proto(dbHandler, grpcServer)

	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln("Failed to close database:", err)
		}
	}(dbHandler.Database())

	fmt.Printf("running at tcp://%v", host)
	return grpcServer.Serve(listen)
}

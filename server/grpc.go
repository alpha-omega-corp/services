package server

import (
	"fmt"
	"github.com/alpha-omega-corp/modules/config"
	"github.com/alpha-omega-corp/modules/database"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewGRPC(host string, c *config.Config, proto func(h *database.Handler, grpc *grpc.Server)) error {
	listen, err := net.Listen("tcp", host)

	if err != nil {
		return err
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

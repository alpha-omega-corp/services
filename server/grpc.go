package server

import (
	"fmt"
	"github.com/alpha-omega-corp/services/database"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewGRPC(host string, h *database.Handler, proto func(db *bun.DB, grpc *grpc.Server)) error {
	listen, err := net.Listen("tcp", host)

	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	proto(h.Database(), grpcServer)

	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(h.Database())

	fmt.Printf("running at tcp://%v", host)
	return grpcServer.Serve(listen)
}

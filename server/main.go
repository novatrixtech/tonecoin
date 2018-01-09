package main

import (
	"context"
	"log"
	"net"

	"github.com/novatrixtech/tonecoin/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8097")
	if err != nil {
		log.Fatalf("Impossivel abrir a porta 8097. Erro: %s\n", err.Error())
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{})
	srv.Serve(listener)
	log.Println("Blockchain server listening at 8097...")
}

//Server represents the implementation of BlockchainServer (gRPC) Server
type Server struct{}

//AddBlock represents the implementation of AddBlock method of BlockchainServer interface
func (s *Server) AddBlock(ctx context.Context, abReq *proto.AddBlockRequest) (abResp *proto.AddBlockResponse, err error) {
	err = nil
	abResp = new(proto.AddBlockResponse)
	return
}

//GetBlockchain represents the implementation of GetBlockchain method of BlockchainServer interface
func (s *Server) GetBlockchain(ctx context.Context, req *proto.GetBlockchainRequest) (resp *proto.GetBlockchainResponse, err error) {
	err = nil
	resp = new(proto.GetBlockchainResponse)
	return
}

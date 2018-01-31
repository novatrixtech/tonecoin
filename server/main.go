package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/novatrixtech/tonecoin/proto"
	"github.com/novatrixtech/tonecoin/server/blockchain"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8097")
	if err != nil {
		log.Fatalf("Impossivel abrir a porta 8097. Erro: %s\n", err.Error())
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	log.Println("Blockchain server listening at 8097...")
	srv.Serve(listener)
}

//Server represents the implementation of BlockchainServer (gRPC) Server
type Server struct {
	Blockchain *blockchain.Blockchain
}

//AddBlock represents the implementation of AddBlock method of BlockchainServer interface
func (s *Server) AddBlock(ctx context.Context, abReq *proto.AddBlockRequest) (abResp *proto.AddBlockResponse, err error) {
	err = nil
	log.Println("[AddBlock] Adding Block...")
	block := s.Blockchain.AddBlock(abReq.Data, abReq.Datatype)
	abResp = &proto.AddBlockResponse{
		Hash: block.Hash,
	}
	log.Println("[AddBlock] Block added...")
	return
}

//GetBlockchain represents the implementation of GetBlockchain method of BlockchainServer interface
func (s *Server) GetBlockchain(ctx context.Context, req *proto.GetBlockchainRequest) (resp *proto.GetBlockchainResponse, err error) {
	err = nil
	log.Println("[GetBlockchain] Getting Block list...")
	resp = new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data.(string),
			Datatype:      proto.Datatype(b.Datatype),
			Timestamp:     b.Timestamp,
		})
	}
	log.Println("[GetBlockchain] Returning Block list...")
	return
}

package main

import (
	"log"
	"net"
	//"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/jacenr/gRPCinDockerDemo/grpcDemo"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetSum(ctx context.Context, in *pb.Req) (*pb.Resp,error) {
	return &pb.Resp{S: in.M + in.N},nil
}

func main() {
        addr := ":8001"
        //if len(os.Args) > 1 {
        //        addr = os.Args[1]
        //}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSumServer(s,&server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

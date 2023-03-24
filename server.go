package main

import (
	"flag"
	"net"

	"github.com/SuhasHebbar/CS739-P2/kvstore"
	pb "github.com/SuhasHebbar/CS739-P2/proto"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {
	slog.Info("Nothing here yet!")

	listenAddress := flag.String("addr", "localhost:8000", "The address the server listens on in the format addr:port. For example localhost: 8000")

	flag.Parse()

	lis, err := net.Listen("tcp", *listenAddress)

	if err != nil {
		slog.Error("Failed to listen on socket", "err", err)
	}

	grpcServer := grpc.NewServer()

	kvStore := kvstore.NewKVRpcServer()

	pb.RegisterRaftRpcServer(grpcServer, kvStore)
	grpcServer.Serve(lis)
}

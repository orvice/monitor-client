package main

import (
	"github.com/orvice/monitor-client/internal/config"
	"github.com/orvice/monitor-client/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func handleGrpc() error {
	lis, err := net.Listen("tcp", config.GrpcAddr)
	if err != nil {
		log.Println("failed to listen grpc server ", err)
		return err
	}
	s := grpc.NewServer()
	monitorClient.RegisterMonitorClientServer(s, newServer())
	return s.Serve(lis)
}

var _ monitorClient.MonitorClientServer = new(Server)

type Server struct {
}

func newServer() *Server {
	return new(Server)
}

// @todo
func (s *Server) Stream(req *monitorClient.StreamRequest, stream monitorClient.MonitorClient_StreamServer) error {
	return nil
}

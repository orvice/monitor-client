package main

import (
	"encoding/json"
	"log"
	"net"

	"github.com/orvice/monitor-client/proto"
	"google.golang.org/grpc"
)

func handleGrpc() error {
	lis, err := net.Listen("tcp", grpcAddr)
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

func (s *Server) Stream(req *monitorClient.StreamRequest, stream monitorClient.MonitorClient_StreamServer) error {
	for {
		res := mtr.GetNetInfo()
		s, err := json.Marshal(res)
		if err != nil {
			continue
		}
		resp := &monitorClient.StreamResponse{
			Body: string(s),
		}
		stream.Send(resp)
	}
	return nil
}

package server

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	client "github.com/coopnorge/interview-backend/src/generated/logistics/api/v1"
)

type GRPCLogisticServer struct {
	service client.CoopLogisticsEngineAPIClient
}

func NewGRPCLogisticServer(service client.CoopLogisticsEngineAPIClient) *GRPCLogisticServer {
	return &GRPCLogisticServer{
		service: service,
	}
}

func (s *GRPCLogisticServer) MoveUnit(ctx context.Context, req *client.MoveUnitRequest) error {
	resp, err := s.service.MoveUnit(ctx, req)
	if err != nil {
		return err
	}

	log := log.New(os.Stdout)
	log.Printf("Message %s", resp)
	return nil
}

func (s *GRPCLogisticServer) UnitReachedWarehouse(ctx context.Context, req *client.UnitReachedWarehouseRequest) error {
	resp, err := s.service.UnitReachedWarehouse(ctx, req)
	if err != nil {
		return err
	}

	log := log.New(os.Stdout)
	log.Printf("Message %s", resp)
	return nil
}

func RunGRPCServer(listenAddr string, service client.CoopLogisticsEngineAPIClient) error {
	GRPCLogisticServer := NewGRPCLogisticServer(service)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []gprc.ServerOptions{}
	server := grpc.NewServer(opts...)
	apiv1.RegisterCoopLogisticsEngineAPIServer(server, GRPCLogisticServer)

	return server.serve(listener)

}

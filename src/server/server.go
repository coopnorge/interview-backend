package server

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/services/client"
	"google.golang.org/grpc"
)

type GRPCLogisticServer struct {
	service client.APILogisticsClient
}

func NewGRPCLogisticServer(service client.APILogisticsClient) *GRPCLogisticServer {
	return &GRPCLogisticServer{
		service: service,
	}
}

func (s *GRPCLogisticServer) MoveUnit(ctx context.Context, req *apiv1.MoveUnitRequest) error {
	resp, err := s.svc.APILogisticsClient.MoveUnit(ctx, req)
	if err != nil {
		return err
	}

	log := log.New(os.Stdout)
	log.Printf("Message %s", resp)
	return nil
}

func (s *GRPCLogisticServer) UnitReachedWarehouse(ctx context.Context, req *apiv1.UnitReachedWarehouseRequest) error {
	resp, err := s.svc.APILogisticsClient.UnitReachedWarehouse(ctx, req)
	if err != nil {
		return err
	}

	log := log.New(os.Stdout)
	log.Printf("Message %s", resp)
	return nil
}

func RunGRPCServer(listenAddr string, service client.APILogisticsClient) error {
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

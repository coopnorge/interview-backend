package server

import (
	"context"
	"net"
	"os"

	"google.golang.org/grpc"

	client "github.com/coopnorge/interview-backend/src/generated/logistics/api/v1"
	serverApi "github.com/coopnorge/interview-backend/src/server/generated/proto"
)

type GRPCLogisticServer struct {
	service client.CoopLogisticsEngineAPIClient
}

func NewGRPCLogisticServer(service client.CoopLogisticsEngineAPIClient) *GRPCLogisticServer {
	return &GRPCLogisticServer{
		service: service,
	}
}

func (s *GRPCLogisticServer) MoveUnit(ctx context.Context, req *client.MoveUnitRequest) (any, error) {
	resp, err := s.service.MoveUnit(ctx, req)
	if err != nil {
		return nil, err
	}

	os.Stdout.Write(resp)
	return &resp, nil
}

func (s *GRPCLogisticServer) UnitReachedWarehouse(ctx context.Context, req *client.UnitReachedWarehouseRequest) (any, error) {
	resp, err := s.service.UnitReachedWarehouse(ctx, req)
	if err != nil {
		return nierr
	}

	os.Stdout.Write(resp)
	return &resp, nil
}

func RunGRPCServer(listenAddr string, service client.CoopLogisticsEngineAPIClient) error {
	GRPCLogisticServer := NewGRPCLogisticServer(service)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []gprc.ServerOptions{}
	server := grpc.NewServer(opts...)
	serverApi.RegisterCoopLogisticsEngineAPIServer(server, GRPCLogisticServer)
	return server.serve(listener)

}

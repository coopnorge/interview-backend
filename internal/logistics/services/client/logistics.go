package client

import (
	"context"
	"strconv"

	apiv1 "github.com/coopnorge/interview-backend/internal/generated/logistics/api/v1"
	"github.com/coopnorge/interview-backend/internal/generated/logistics/api/v1/openapi"
	"github.com/coopnorge/interview-backend/internal/logistics/config"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServiceSetForClient providers
var ServiceSetForClient = wire.NewSet(NewLogisticsClient)

// TransportType represents how messages must be sent.
type TransportType byte

const (
	// TransportTypeGRPC default mode how client will send messages to server.
	TransportTypeGRPC TransportType = iota
	// TransportTypeHTTP generic HTTP requests.
	TransportTypeHTTP
)

const (
	TransportTypeGRPCStr = "gRPC"
	TransportTypeHTTPStr = "HTTP"
)

// APILogisticsClient to send requests about cargo unit movements
type APILogisticsClient struct {
	apiClientGRPC apiv1.CoopLogisticsEngineAPIClient
	apiClientHTTP *openapi.APIClient

	transportType   TransportType
	transportScheme string
	conn            *grpc.ClientConn
}

// NewLogisticsClient instance
func NewLogisticsClient(cfg *config.ClientAppConfig) *APILogisticsClient {
	return &APILogisticsClient{
		transportType:   parseTransportType(cfg.TransportTypeProtocol),
		transportScheme: cfg.Scheme,
	}
}

// Connect to gRPC API
func (lc *APILogisticsClient) Connect(serverAddr string, ctx context.Context) error {
	if lc.transportType.Is(TransportTypeGRPC) {
		conn, dialErr := grpc.DialContext(
			ctx,
			serverAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if dialErr != nil {
			return dialErr
		}

		lc.conn = conn
		lc.apiClientGRPC = apiv1.NewCoopLogisticsEngineAPIClient(lc.conn)

		return nil
	}

	httpConfig := openapi.NewConfiguration()
	httpConfig.Host = serverAddr
	httpConfig.Scheme = lc.transportScheme
	lc.apiClientHTTP = openapi.NewAPIClient(httpConfig)

	return nil
}

// Disconnect from gRPC API
func (lc *APILogisticsClient) Disconnect() error {
	return lc.conn.Close()
}

// MoveUnit to new location
func (lc *APILogisticsClient) MoveUnit(ctx context.Context, req *apiv1.MoveUnitRequest) (responseErr error) {
	if lc.transportType.Is(TransportTypeGRPC) {
		_, responseErr = lc.apiClientGRPC.MoveUnit(ctx, req)
		return
	}

	_, _, responseErr = lc.apiClientHTTP.CoopLogisticsEngineAPIAPI.
		CoopLogisticsEngineAPIMoveUnit(ctx).
		CargoUnitId(strconv.FormatInt(req.GetCargoUnitId(), 10)).
		LocationLatitude(int64(req.GetLocation().GetLatitude())).
		LocationLongitude(int64(req.GetLocation().GetLongitude())).
		Execute()

	return
}

// UnitReachedWarehouse report that reach warehouse
func (lc *APILogisticsClient) UnitReachedWarehouse(ctx context.Context, req *apiv1.UnitReachedWarehouseRequest) (responseErr error) {
	if lc.transportType.Is(TransportTypeGRPC) {
		_, responseErr = lc.apiClientGRPC.UnitReachedWarehouse(ctx, req)
		return
	}

	_, _, responseErr = lc.apiClientHTTP.CoopLogisticsEngineAPIAPI.
		CoopLogisticsEngineAPIUnitReachedWarehouse(ctx).
		AnnouncementCargoUnitId(strconv.FormatInt(req.GetAnnouncement().GetCargoUnitId(), 10)).
		AnnouncementWarehouseId(strconv.FormatInt(req.GetAnnouncement().GetWarehouseId(), 10)).
		AnnouncementMessage(req.GetAnnouncement().GetMessage()).
		LocationLatitude(int64(req.GetLocation().GetLatitude())).
		LocationLongitude(int64(req.GetLocation().GetLongitude())).
		Execute()

	return
}

// Is TransportType matching needed one.
func (tt *TransportType) Is(needed TransportType) bool {
	return tt != nil && *tt == needed
}

// parseTransportType allows parse string and get supported TransportType or fallback to default.
func parseTransportType(transportType string) TransportType {
	switch transportType {
	case TransportTypeHTTPStr:
		return TransportTypeHTTP
	default:
		return TransportTypeGRPC
	}
}

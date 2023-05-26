package logistics

import (
    "context"
    "errors"
    "fmt"
    "log"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"

    api "github.com/coopnorge/interview-backend/internal/app/logistics/api/v1"
    "github.com/coopnorge/interview-backend/internal/app/logistics/model"
    "github.com/coopnorge/interview-backend/internal/app/logistics/services/client"
    "github.com/coopnorge/interview-backend/internal/app/logistics/services/operator"
    "github.com/coopnorge/interview-backend/internal/app/pkg/generator"
)

const (
    appName    = "Coop Logistics Engine"
    apiAddress = "127.0.0.1:50051" // TODO Improve later to use CMD ARGs

    maxWarehouses = 1<<8 - 1
    maxCargoUnits = 1 << 10
)

// ServiceInstance of application
type ServiceInstance struct {
    ctx       context.Context
    ctxCancel context.CancelFunc

    logisticsClient *client.APILogisticsClient
    worldOperator   *operator.WorldOperator
}

// NewServiceInstance constructor
func NewServiceInstance(lc *client.APILogisticsClient, wo *operator.WorldOperator) (*ServiceInstance, error) {
    log.Printf("%s, initializing...\n", appName)

    serviceCtx, serviceCtxCancel := context.WithCancel(context.Background())
    connCtx, connCtxCancel := context.WithTimeout(serviceCtx, 10*time.Second)
    defer connCtxCancel()

    if connErr := lc.Connect(apiAddress, connCtx); connErr != nil {
        serviceCtxCancel()
        err := errors.New(fmt.Sprintf(
            "%s, failed to connect to API (%s), error: %v",
            appName,
            apiAddress,
            connErr,
        ))

        return nil, err
    }

    service := &ServiceInstance{
        ctx:       serviceCtx,
        ctxCancel: serviceCtxCancel,

        logisticsClient: lc,
        worldOperator:   wo,
    }

    worldPopulationErr := wo.Populate(
        uint32(rand.Intn(maxWarehouses-10+1)+10),
        uint32(rand.Intn(maxCargoUnits-10+1)+10),
    )
    if worldPopulationErr != nil {
        return nil, worldPopulationErr
    }

    return service, nil
}

// Run app
func (s *ServiceInstance) Run() error {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

    go func() { // Handle graceful shutdown
        <-signals // Wait for the signal

        log.Printf("%s, shutting down...\n", appName)

        s.ctxCancel()
        if s.logisticsClient != nil {
            _ = s.logisticsClient.Disconnect()
        }

        log.Printf("%s, stopped!\n", appName)

        os.Exit(0)
    }()

    deliveryUnits := s.worldOperator.GetDeliveryUnit()
    totalDeliveryUnits := len(deliveryUnits)

    for {
        var wg sync.WaitGroup
        unitsReachedObjective := 0

        // Check if all units reached goal
        for _, unit := range deliveryUnits {
            if unit.Metadata == true {
                unitsReachedObjective++
            }
        }

        if unitsReachedObjective == totalDeliveryUnits {
            log.Println("All delivery units reached warehouse...")
            break
        }

        for _, unit := range deliveryUnits {
            if unit.Metadata == true {
                continue
            }

            wg.Add(1)
            go s.processDelivery(unit, &wg)

        }

        wg.Wait()
    }

    return nil
}

func (s *ServiceInstance) processDelivery(unit *model.GraphNode, wg *sync.WaitGroup) {
    defer wg.Done()

    time.Sleep(30 * time.Millisecond)

    oldCoordinate := *unit.Coordinate
    newCoordinate := s.worldOperator.MoveDeliveryUniToNearestWarehouse(unit.ID)
    unitMessage := fmt.Sprintf("%s moving to - X:%d, Y:%d", unit.Name, newCoordinate.X, newCoordinate.Y)

    log.Println(unitMessage)

    moveErr := s.logisticsClient.MoveUnit(
        s.ctx,
        &api.MoveUnitRequest{
            CargoUnitId: int64(unit.ID),
            Location: &api.Location{
                X: int64(newCoordinate.X),
                Y: int64(newCoordinate.Y),
            },
        },
    )
    if moveErr != nil {
        log.Printf("filed to send MoveUnit %s, API error: %v\n", unitMessage, moveErr)
        return
    } else if newCoordinate != oldCoordinate {
        return
    }

    announcement := fmt.Sprintf("%s - Reached Objective.", unitMessage)
    warehouse := s.worldOperator.FindEntityByCoordinate(newCoordinate, generator.Warehouses)
    if warehouse == nil {
        log.Printf("Warehouses not found in coordinates X:%d Y:%d", newCoordinate.X, newCoordinate.Y)
        return
    }

    reachErr := s.logisticsClient.UnitReachedWarehouse(
        s.ctx,
        &api.UnitReachedWarehouseRequest{
            Location: &api.Location{X: int64(newCoordinate.X), Y: int64(newCoordinate.Y)},
            Announcement: &api.WarehouseAnnouncement{
                CargoUnitId: int64(unit.ID),
                WarehouseId: int64(warehouse.ID),
                Message:     announcement,
            },
        },
    )
    if reachErr != nil {
        log.Printf("filed to send UnitReachedWarehouse %s, API error: %v\n", unitMessage, moveErr)
        return
    }

    log.Println(announcement)
    unit.Metadata = true // Unit reached Warehouse

    return
}

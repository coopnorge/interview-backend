package generator

import (
    "fmt"

    "github.com/brianvoe/gofakeit/v6"
    "github.com/coopnorge/interview-backend/internal/app/logistics/model"
)

// ActorType kind
type ActorType byte

const (
    Warehouses ActorType = iota
    CargoUnits
)

// AddNewActors by type to the model.Graph with actorNumber and from what ID it must be added (idPrefix)
func AddNewActors(t ActorType, g *model.Graph, actorNumber uint, idPrefix uint) {
    locsAndRange := int(actorNumber)
    locations := NewCoordinates(locsAndRange, 1<<8-1, 1<<8-1)
    for i := uint(0); i < actorNumber; i++ {
        actorNode := model.GraphNode{ID: idPrefix + i}

        switch t {
        case Warehouses:
            actorNode.Name = fmt.Sprintf("Warehouse: %s - %s", gofakeit.City(), gofakeit.Company())
            actorNode.Type = Warehouses
        case CargoUnits:
            actorNode.Name = fmt.Sprintf("CargoUnit: %s - %s", gofakeit.CarMaker(), gofakeit.CarModel())
            actorNode.Type = CargoUnits
            actorNode.Metadata = false // TODO Later can be fixed, but used to indicate if unit reached objective
        }

        actorNode.Coordinate = &locations[i]

        g.AddNode(actorNode)
    }
}

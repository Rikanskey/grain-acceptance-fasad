package domain

import "time"

type ProcessType string

const (
	acceptance ProcessType = "Acceptance"
	shipment   ProcessType = "Shipment"
)

func (pt ProcessType) String() string {
	switch pt {
	case acceptance:
		return "Acceptance"
	case shipment:
		return "Shipment"
	}

	return ""
}

type GrainCard struct {
	id       uint
	customer Company
	sender   Company
	date     time.Time
	//consignments []Consignment
	storage uint
	// Enum string type: acceptance or shipment
	processType ProcessType
}

func (gc *GrainCard) Id() uint {
	return gc.id
}

func (gc *GrainCard) Customer() Company {
	return gc.customer
}

func (gc *GrainCard) Sender() Company {
	return gc.sender
}

func (gc *GrainCard) Date() time.Time {
	return gc.date
}

func (gc *GrainCard) ProcessType() ProcessType {
	return gc.processType
}

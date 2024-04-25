package types

type (
	ParkingStore interface {
		AddParking(p *ParkingAddPayload) error
		GetAllFreeSlots(parkingName string) ([]EmptySlotPayload, error)
		GetAllFullSlots(parkingName string) ([]FullSlotPayload, error)
	}

	VehicleStore interface {
		AddVehicle(v *VehiclePayload) (*VehiclePayloadReturn, error)
		RemoveVehicle(wincode string) (*VehiclePayloadReturn, error)
		GetVehicle(wincode string) (*VehiclePayloadReturn, error)
	}
)

type (
	ParkingAddPayload struct {
		Name   string `json:"name"`
		Region string `json:"region"`
		Max    int    `json:"max"`
	}
)

type (
	VehiclePayload struct {
		WinCode     string `json:"wincode"`
		ParkingName string `json:"parking"`
	}

	VehiclePayloadReturn struct {
		ParkingName string `json:"parking"`
		BlockName   string `json:"block"`
		SlotNumber  int    `json:"slot"`
	}
)

type (
	SlotPayload struct {
		ParkingName string `json:"parking"`
	}

	EmptySlotPayload struct {
		BlockName  string `json:"block"`
		SlotNumber int    `json:"slot"`
	}

	FullSlotPayload struct {
		BlockName  string `json:"block"`
		SlotNumber int    `json:"slot"`
		WinCode    string `json:"wincode"`
	}
)

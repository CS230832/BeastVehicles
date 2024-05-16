package types

type (
	AdminStore interface {
		AddAdmin(admin *AdminPayload) error
		RemoveAdmin(email string) error
		GetAdmin(email string) (*AdminPayload, error)

		AddToken(email string, content string) error
		RemoveToken(email string, content string) error
		GetTokens(email string) ([]string, error)
	}

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

	BlockStore interface {
		GetFreeSlots(parking string, block string) ([]EmptySlotPayload, error)
		GetFullSlots(parking string, block string) ([]FullSlotPayload, error)
	}
)

type (
	AdminPayload struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		IsSuper   bool   `json:"is_super"`
	}

	AdminLoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

	EmptySlotReturnPayload struct {
		SlotNumber int `json:"slot"`
	}

	FullSlotReturnPayload struct {
		SlotNumber int    `json:"slot"`
		WinCode    string `json:"wincode"`
	}
)

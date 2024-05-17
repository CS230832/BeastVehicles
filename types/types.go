package types

type UserRole string

const (
	Root    UserRole = "root"
	Manager UserRole = "manager"
	Admin   UserRole = "admin"
)

type (
	UserStore interface {
		AddUser(user *UserRegisterPayload) error
		RemoveUser(username string) error
		GetUser(username string) (*UserPayload, error)

		AddLoginToken(username string, token string) error
		RemoveLoginToken(username string, token string) error
		HasLoginToken(username string, token string) (bool, error)
		RemoveAllLoginTokens(username string) error
	}

	ParkingStore interface {
		AddParking(parking *ParkingPayload) error
		RemoveParking(name string) error
		GetParking(name string) (*ParkingPayload, error)
	}

	BlockStore interface {
		GetBlock(name string, parking string) (*BlockPayload, error)
		GetBlockFreeSlots(name string, parking string) (*BlockPayload, error)
		GetBlockFullSlots(name string, parking string) (*BlockPayload, error)

		GetBlocks(parking string) (*BlockSetPayload, error)
		GetFreeSlotsInBlocks(parking string) (*BlockSetPayload, error)
		GetFullSlotsInBlocks(parking string) (*BlockSetPayload, error)
	}

	VehicleStore interface {
		AddVehicle(payload *VehicleRegisterPayload, parking string) (*VehiclePayload, error)
		RemoveVehicle(wincode string) error
		GetVehicle(wincode string) (*VehiclePayload, error)

		AddVehicles(set *VehicleSetRegisterPayload, parking string) ([]VehiclePayload, error)
		RemoveVehicles(wincodes []string) error
		GetVehicles(wincodes []string) ([]VehiclePayload, error)
	}
)

type (
	UserRegisterPayload struct {
		UserName    string   `json:"username"`
		Password    string   `json:"password"`
		Role        UserRole `json:"role"`
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		ParkingName string   `json:"parking"`
	}

	UserLoginPayload struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	UserPayload struct {
		ID          int      `json:"-"`
		UserName    string   `json:"username,omitempty"`
		Password    string   `json:"-"`
		Role        UserRole `json:"role,omitempty"`
		FirstName   string   `json:"first_name,omitempty"`
		LastName    string   `json:"last_name,omitempty"`
		ParkingName string   `json:"parking,omitempty"`
	}
)

type (
	ParkingPayload struct {
		ID       int    `json:"-"`
		Name     string `json:"name"`
		Capacity int    `json:"capacity"`
		Region   string `json:"region,omitempty"`
	}
)

type (
	BlockPayload struct {
		Name  string
		Slots []SlotPayload
	}

	SlotPayload struct {
		Number  int     `json:"slot"`
		WinCode *string `json:"wincode,omitempty"`
	}

	BlockSetPayload struct {
		Blocks map[string][]SlotPayload
	}
)

type (
	VehicleRegisterPayload struct {
		WinCode string `json:"wincode"`
	}

	VehicleSetRegisterPayload struct {
		Set []string `json:"wincodes"`
	}

	VehiclePayload struct {
		ID      int    `json:"-"`
		WinCode string `json:"wincode,omitempty"`
		Parking string `json:"parking,omitempty"`
		Block   string `json:"block,omitempty"`
		Slot    string `json:"slot,omitempty"`
	}
)

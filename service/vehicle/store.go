package vehicle

import (
	"CS230832/BeastVehicles/types"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddVehicle(v *types.VehiclePayload) (*types.VehiclePayloadReturn, error) {
	var result types.VehiclePayloadReturn

	row := s.db.QueryRow("SELECT block_name, slot_number FROM AddVehicle($1, $2)", v.WinCode, v.ParkingName)

	if row == nil {
		return &result, fmt.Errorf("vehicle with wincode %s already exists", v.WinCode)
	}

	if err := row.Scan(&result.BlockName, &result.SlotNumber); err != nil {
		return &result, err
	}

	result.ParkingName = v.ParkingName

	return &result, nil
}

func (s *Store) RemoveVehicle(wincode string) (*types.VehiclePayloadReturn, error) {
	var result types.VehiclePayloadReturn

	row := s.db.QueryRow("SELECT parking_name, block_name, slot_number FROM RemoveVehicle($1)", wincode)

	if row == nil {
		return &result, fmt.Errorf("vehicle with wincode %s does not exist", wincode)
	}

	if err := row.Scan(&result.ParkingName, &result.BlockName, &result.SlotNumber); err != nil {
		return &result, err
	}

	return &result, nil
}

func (s *Store) GetVehicle(wincode string) (*types.VehiclePayloadReturn, error) {
	var result types.VehiclePayloadReturn

	row := s.db.QueryRow("SELECT parking_name, block_name, slot_number FROM FindVehicle($1)", wincode)

	if row == nil {
		return &result, fmt.Errorf("vehicle with wincode %s does not exist", wincode)
	}

	if err := row.Scan(&result.ParkingName, &result.BlockName, &result.SlotNumber); err != nil {
		return &result, err
	}

	return &result, nil
}
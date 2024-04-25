package parking

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

func (s *Store) AddParking(p *types.ParkingAddPayload) error {
	var id int

	row := s.db.QueryRow("SELECT AddParking($1, $2, $3)", p.Name, p.Region, p.Max)

	if row == nil {
		return fmt.Errorf("parking with name %s already exists", p.Name)
	}

	return row.Scan(&id)
}

func (s *Store) GetAllFreeSlots(parkingName string) ([]types.EmptySlotPayload, error) {
	var slots []types.EmptySlotPayload

	rows, err := s.db.Query("SELECT block_name, slot_number FROM FindAllFreeSlots($1)", parkingName)

	if err != nil {
		return slots, err
	}

	defer rows.Close()

	for rows.Next() {
		var slot types.EmptySlotPayload
		if err := rows.Scan(&slot.BlockName, &slot.SlotNumber); err != nil {
			return slots, err
		}
		slots = append(slots, slot)
	}

	return slots, nil
}

func (s *Store) GetAllFullSlots(parkingName string) ([]types.FullSlotPayload, error) {
	var slots []types.FullSlotPayload

	rows, err := s.db.Query("SELECT block_name, slot_number, wincode FROM FindAllFullSlots($1)", parkingName)

	if err != nil {
		return slots, err
	}

	defer rows.Close()

	for rows.Next() {
		var slot types.FullSlotPayload
		if err := rows.Scan(&slot.BlockName, &slot.SlotNumber, &slot.WinCode); err != nil {
			return slots, err
		}
		slots = append(slots, slot)
	}

	return slots, nil
}

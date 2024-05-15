package block

import (
	"CS230832/BeastVehicles/types"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetFreeSlots(parking string, block string) ([]types.EmptySlotPayload, error) {
	var slots []types.EmptySlotPayload

	rows, err := s.db.Query("SELECT slot_number FROM BlockFindAllFreeSlots($1, $2);", parking, block)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var slot types.EmptySlotPayload
		slot.BlockName = block
		if err := rows.Scan(&slot.SlotNumber); err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}

	return slots, nil
}

func (s *Store) GetFullSlots(parking string, block string) ([]types.FullSlotPayload, error) {
	var slots []types.FullSlotPayload

	rows, err := s.db.Query("SELECT slot_number, wincode FROM BlockFindAllFullSlots($1, $2);", parking, block)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var slot types.FullSlotPayload
		slot.BlockName = block
		if err := rows.Scan(&slot.SlotNumber, &slot.WinCode); err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}

	return slots, nil
}
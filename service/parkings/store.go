package parkings

import (
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddParking(parking *types.ParkingPayload) error {
	if parking.Name == "" {
		return fmt.Errorf("no parking name given")
	}

	if parking.Capacity == 0 {
		return fmt.Errorf("parking capacity cannot be 0")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var exists bool
	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM Parkings WHERE name = $1);",
		parking.Name,
	).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("parking with name '%s' already exists", parking.Name)
	}

	var parkingID int
	if err := tx.QueryRow(
		"INSERT INTO Parkings (name, capacity, region) VALUES ($1, $2, $3) RETURNING id;",
		parking.Name,
		parking.Capacity,
		parking.Region,
	).Scan(&parkingID); err != nil {
		return err
	}

	blockCount := parking.Capacity / 50
	remainingSlots := parking.Capacity % 50

	for i := 0; i < blockCount; i++ {
		if err := s.addBlock(tx, i, 50, parkingID); err != nil {
			return err
		}
	}

	if remainingSlots > 0 {
		if err := s.addBlock(tx, blockCount, remainingSlots, parkingID); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *Store) RemoveParking(name string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	parkingID, err := s.getParkingID(tx, name)
	if err != nil {
		return err
	}

	if _, err := tx.Exec("DELETE FROM Vehicles WHERE parking_id = $1;", parkingID); err != nil {
		return err
	}

	rows, err := tx.Query("SELECT id FROM Blocks WHERE parking_id = $1;", parkingID)
	if err != nil {
		return err
	}

	var blockIDs []int
	for rows.Next() {
		var blockID int
		if err := rows.Scan(&blockID); err != nil {
			return err
		}
		blockIDs = append(blockIDs, blockID)
	}

	for _, blockID := range blockIDs {
		if _, err := tx.Exec("DELETE FROM Slots WHERE block_id = $1;", blockID); err != nil {
			return err
		}

		if _, err := tx.Exec("DELETE FROM Blocks WHERE id = $1;", blockID); err != nil {
			return err
		}
	}

	if _, err := tx.Exec("DELETE FROM Parkings WHERE id = $1;", parkingID); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) GetParking(name string) (*types.ParkingPayload, error) {
	var parking types.ParkingPayload

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if err := tx.QueryRow(
		"SELECT id, name, capacity, region FROM Parkings WHERE name = $1;",
		name,
	).Scan(
		&parking.ID,
		&parking.Name,
		&parking.Capacity,
		&parking.Region,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no parking found with name '%s'", name)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &parking, nil
}

func (s *Store) addBlock(tx *sql.Tx, blockNumber int, slotCount int, parkingID int) error {
	blockName := utils.NumToBlockName(blockNumber)

	var blockID int
	if err := tx.QueryRow(
		"INSERT INTO Blocks (name, parking_id) VALUES ($1, $2) RETURNING id;",
		blockName,
		parkingID,
	).Scan(&blockID); err != nil {
		return err
	}

	for i := 0; i < slotCount; i++ {
		if err := s.addSlot(tx, i+1, blockID, parkingID); err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) addSlot(tx *sql.Tx, slot_number int, blockID int, parkingID int) error {
	if _, err := tx.Exec(
		"INSERT INTO Slots (number, parking_id, block_id) VALUES ($1, $2, $3);",
		slot_number,
		parkingID,
		blockID,
	); err != nil {
		return err
	}

	return nil
}

func (s *Store) getParkingID(tx *sql.Tx, name string) (int, error) {
	var parkingID int

	if err := tx.QueryRow(
		"SELECT id FROM Parkings WHERE name = $1;",
		name,
	).Scan(&parkingID); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no parking found with name '%s'", name)
		}
		return 0, err
	}

	return parkingID, nil
}

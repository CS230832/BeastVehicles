package blocks

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

func (s *Store) GetBlock(name string, parking string) (*types.BlockPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	exists, err := s.doesParkingExist(tx, parking)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no parking found with name '%s'", parking)
	}

	block, err := s.getBlock(tx, name, parking)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return block, nil
}

func (s *Store) GetBlockFreeSlots(name string, parking string) (*types.BlockPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	exists, err := s.doesParkingExist(tx, parking)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no parking found with name '%s'", parking)
	}

	block, err := s.getBlockFreeSlots(tx, name, parking)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return block, nil
}

func (s *Store) GetBlockFullSlots(name string, parking string) (*types.BlockPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	exists, err := s.doesParkingExist(tx, parking)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no parking found with name '%s'", parking)
	}

	block, err := s.getBlockFullSlots(tx, name, parking)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return block, nil
}

func (s *Store) GetBlocks(parking string) (*types.BlockSetPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	exists, err := s.doesParkingExist(tx, parking)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no parking found with name '%s'", parking)
	}

	rows, err := tx.Query(
		`SELECT b.name AS block_name
		FROM Blocks b
		JOIN Parkings p ON b.parking_id = p.id
		WHERE p.name = $1
		ORDER BY b.name;`,
		parking,
	)
	if err != nil {
		return nil, err
	}
	
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	set := types.BlockSetPayload{Blocks: make(map[string][]types.SlotPayload)}
	for _, name := range names {
		block, err := s.getBlock(tx, name, parking)
		if err != nil {
			return nil, err
		}
		set.Blocks[name] = block.Slots
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &set, nil
}

func (s *Store) GetFreeSlotsInBlocks(parking string) (*types.BlockSetPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	exists, err := s.doesParkingExist(tx, parking)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no parking found with name '%s'", parking)
	}

	rows, err := tx.Query(
		`SELECT b.name AS block_name
		FROM Blocks b
		JOIN Parkings p ON b.parking_id = p.id
		WHERE p.name = $1
		ORDER BY b.name;`,
		parking,
	)
	if err != nil {
		return nil, err
	}
	
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	set := types.BlockSetPayload{Blocks: make(map[string][]types.SlotPayload)}
	for _, name := range names {
		block, err := s.getBlockFreeSlots(tx, name, parking)
		if err != nil {
			return nil, err
		}
		set.Blocks[name] = block.Slots
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &set, nil
}

func (s *Store) GetFullSlotsInBlocks(parking string) (*types.BlockSetPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	exists, err := s.doesParkingExist(tx, parking)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("no parking found with name '%s'", parking)
	}

	rows, err := tx.Query(
		`SELECT b.name AS block_name
		FROM Blocks b
		JOIN Parkings p ON b.parking_id = p.id
		WHERE p.name = $1
		ORDER BY b.name;`,
		parking,
	)
	if err != nil {
		return nil, err
	}
	
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	set := types.BlockSetPayload{Blocks: make(map[string][]types.SlotPayload)}
	for _, name := range names {
		block, err := s.getBlockFullSlots(tx, name, parking)
		if err != nil {
			return nil, err
		}
		set.Blocks[name] = block.Slots
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &set, nil
}

func (s *Store) getBlock(tx *sql.Tx, name string, parking string) (*types.BlockPayload, error) {
	if parking == "" || name == "" {
		return nil, fmt.Errorf("neither name nor parking should be empty strings")
	}

	rows, err := tx.Query(
		`SELECT 
			s.number AS slot_number,
			v.wincode
		FROM Slots s 
		JOIN Blocks b ON s.block_id = b.id 
		JOIN Parkings p ON b.parking_id = p.id 
		LEFT JOIN Vehicles v ON s.id = v.slot_id 
		WHERE p.name = $1 AND b.name = $2
		ORDER BY s.number;`,
		parking,
		name,
	)
	if err != nil {
		return nil, err
	}

	block := types.BlockPayload{Name: name}
	var slot types.SlotPayload
	for rows.Next() {
		if err := rows.Scan(&slot.Number, &slot.WinCode); err != nil {
			return nil, err
		}
		block.Slots = append(block.Slots, slot)
	}
	return &block, nil
}

func (s *Store) getBlockFreeSlots(tx *sql.Tx, name string, parking string) (*types.BlockPayload, error) {
	if parking == "" || name == "" {
		return nil, fmt.Errorf("neither name nor parking should be empty strings")
	}

	rows, err := tx.Query(
		`SELECT s.number AS slot_number
		FROM Slots s 
		JOIN Blocks b ON s.block_id = b.id 
		JOIN Parkings p ON b.parking_id = p.id
		WHERE p.name = $1 AND b.name = $2 AND s.is_empty = TRUE
		ORDER BY s.number;`,
		parking,
		name,
	)
	if err != nil {
		return nil, err
	}

	block := types.BlockPayload{Name: name}
	var slot types.SlotPayload
	for rows.Next() {
		if err := rows.Scan(&slot.Number); err != nil {
			return nil, err
		}
		block.Slots = append(block.Slots, slot)
	}
	return &block, nil
}

func (s *Store) getBlockFullSlots(tx *sql.Tx, name string, parking string) (*types.BlockPayload, error) {
	if parking == "" || name == "" {
		return nil, fmt.Errorf("neither name nor parking should be empty strings")
	}

	rows, err := tx.Query(
		`SELECT 
			s.number AS slot_number,
			v.wincode
		FROM Slots s 
		JOIN Blocks b ON s.block_id = b.id 
		JOIN Parkings p ON b.parking_id = p.id 
		JOIN Vehicles v ON s.id = v.slot_id 
		WHERE p.name = $1 AND b.name = $2
		ORDER BY s.number;`,
		parking,
		name,
	)
	if err != nil {
		return nil, err
	}

	block := types.BlockPayload{Name: name}
	var slot types.SlotPayload
	for rows.Next() {
		if err := rows.Scan(&slot.Number, &slot.WinCode); err != nil {
			return nil, err
		}
		block.Slots = append(block.Slots, slot)
	}
	return &block, nil
}

func (s *Store) doesParkingExist(tx *sql.Tx, parking string) (bool, error) {
	var exists bool

	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM Parkings WHERE name = $1);",
		parking,
	).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("no parking found with name '%s'", parking)
		}
		return false, err
	}

	return exists, nil
}

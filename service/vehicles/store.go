package vehicles

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

func (s *Store) AddVehicle(payload *types.VehicleRegisterPayload, parking string) (*types.VehiclePayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	vehicle, err := s.addVehicle(tx, payload.WinCode, parking)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Store) RemoveVehicle(wincode string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := s.removeVehicle(tx, wincode); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) GetVehicle(wincode string) (*types.VehiclePayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	vehicle, err := s.getVehicle(tx, wincode)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Store) AddVehicles(set *types.VehicleSetRegisterPayload, parking string) ([]types.VehiclePayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var vehicles []types.VehiclePayload
	for _, wincode := range set.Set {
		vehicle, err := s.addVehicle(tx, wincode, parking)
		if err != nil {
			return nil, err
		}
		vehicle.WinCode = wincode
		vehicles = append(vehicles, *vehicle)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (s *Store) RemoveVehicles(wincodes []string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, wincode := range wincodes {
		if err := s.removeVehicle(tx, wincode); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *Store) GetVehicles(wincodes []string) ([]types.VehiclePayload, error ){
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var vehicles []types.VehiclePayload

	for _, wincode := range wincodes {
		vehicle, err := s.getVehicle(tx, wincode)
		if err != nil {
			return nil, err
		}
		vehicle.WinCode = wincode
		vehicles = append(vehicles, *vehicle)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return vehicles, err
}

func (s *Store) addVehicle(tx *sql.Tx, wincode string, parking string) (*types.VehiclePayload, error) {
	if wincode == "" {
		return nil, fmt.Errorf("wincode of vehicle cannot be empty")
	}

	var exists bool
	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM Vehicles WHERE wincode = $1);",
		wincode,
	).Scan(&exists); err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("vehicle with wincode '%s' already exists", wincode)
	}

	vehicle := types.VehiclePayload{Parking: parking}
	var slotID int

	if err := tx.QueryRow(
		`WITH next_free_slot AS (
			SELECT s.id AS slot_id, b.id AS block_id
			FROM Slots s
			JOIN Blocks b ON s.block_id = b.id
			WHERE s.is_empty = TRUE AND s.parking_id = (
				SELECT id FROM Parkings WHERE name = $1
			)
			ORDER BY b.id, s.number
			LIMIT 1
		)
		INSERT INTO Vehicles (wincode, parking_id, block_id, slot_id)
		SELECT $2, (
			SELECT id FROM Parkings WHERE name = $1
		), block_id, slot_id
		FROM next_free_slot
		RETURNING id, slot_id;`,
		parking,
		wincode,
	).Scan(&vehicle.ID, &slotID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no free space left in parking '%s'", parking)
		}
		return nil, err
	}

	if _, err := tx.Exec(
		"UPDATE Slots SET is_empty = FALSE WHERE id = $1;",
		slotID,
	); err != nil {
		return nil, err
	}

	if err := tx.QueryRow(
		`SELECT b.name, s.number
		FROM Vehicles v
		JOIN Blocks b ON v.block_id = b.id
		JOIN Slots s ON v.slot_id = s.id
		WHERE v.id = $1;`,
		vehicle.ID,
	).Scan(&vehicle.Block, &vehicle.Slot); err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (s *Store) removeVehicle(tx *sql.Tx, wincode string) error {
	var slotID int

	if err := tx.QueryRow(
		"DELETE FROM Vehicles WHERE wincode = $1 RETURNING slot_id;",
		wincode,
	).Scan(&slotID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no vehicle found with wincode '%s'", wincode)
		}
	}

	if _, err := tx.Exec(
		"UPDATE Slots SET is_empty = TRUE WHERE id = $1;",
		slotID,
	); err != nil {
		return err
	}

	return nil
}

func (s *Store) getVehicle(tx *sql.Tx, wincode string) (*types.VehiclePayload, error) {
	var vehicle types.VehiclePayload

	if err := tx.QueryRow(
		`SELECT v.id, p.name, b.name, s.number
		FROM Vehicles v
		JOIN Parkings p ON v.parking_id = p.id
		JOIN Blocks b ON v.block_id = b.id
		JOIN Slots s ON v.slot_id = s.id
		WHERE v.wincode = $1;`,
		wincode,
	).Scan(
		&vehicle.ID,
		&vehicle.Parking,
		&vehicle.Block,
		&vehicle.Slot,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no vehicle found with wincode '%s'", wincode)
		}
		return nil, err
	}

	return &vehicle, nil
}

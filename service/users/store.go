package users

import (
	"CS230832/BeastVehicles/service/auth"
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

func (s *Store) AddUser(user *types.UserRegisterPayload) error {
	if user.Role != types.Root &&
		user.Role != types.CEO &&
		user.Role != types.Manager {
		return fmt.Errorf("invalid role")
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var exists bool
	if err := tx.QueryRow("SELECT EXISTS(SELECT 1 FROM Users WHERE username = $1);",
		user.UserName,
	).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("user with username '%s' already exists", user.UserName)
	}

	if _, err := tx.Exec(
		"INSERT INTO Users (username, password, role, first_name, last_name) VALUES ($1, $2, $3, $4, $5);",
		user.UserName,
		hashedPassword,
		user.Role,
		user.FirstName,
		user.LastName,
	); err != nil {
		return err
	}

	if !(user.Role == types.Root && user.ParkingName == "") {
		if user.ParkingName == "" {
			return fmt.Errorf("no parking name given")
		}

		var parkingID int

		if err := s.db.QueryRow(
			"SELECT id FROM Parkings WHERE name = $1;",
			user.ParkingName,
		).Scan(&parkingID); err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("no parking found with name '%s'", user.ParkingName)
			}
			return err
		}

		if _, err := tx.Exec(
			"UPDATE Users SET parking_id = $1 WHERE username = $2;",
			parkingID,
			user.UserName,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *Store) RemoveUser(username string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var exists bool
	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM Users WHERE username = $1);",
		username,
	).Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("no user found with username '%s'", username)
	}

	if _, err := tx.Exec("DELETE FROM Users WHERE username = $1;", username); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) GetUserByUserName(username string) (*types.UserPayload, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var parkingID int
	user := types.UserPayload{UserName: username}

	if err := tx.QueryRow(
		"SELECT id, password, role, first_name, last_name FROM Users WHERE username = $1;",
		username,
	).Scan(
		&user.ID,
		&user.Password,
		&user.Role,
		&user.FirstName,
		&user.LastName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with username '%s'", username)
		}
		return nil, err
	}

	if err := tx.QueryRow(
		"SELECT parking_id FROM Users WHERE id = $1;",
		user.ID,
	).Scan(&parkingID); err != nil {
		if user.Role == types.Root {
			if err := tx.Commit(); err != nil {
				return nil, err
			}

			return &user, nil
		}

		return nil, err
	}

	if err := tx.QueryRow(
		"SELECT name FROM Parkings WHERE id = $1;",
		parkingID,
	).Scan(&user.ParkingName); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no parking found with id '%d'", parkingID)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) AddLoginToken(username string, token string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userID, err := getUserID(tx, username)
	if err != nil {
		return err
	}

	var exists bool
	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM LoginTokens WHERE user_id = $1 AND content = $2);",
		userID,
		token,
	).Scan(&exists); err != nil {
		return err
	}

	if exists {
		tx.Rollback()
		return nil
	}

	if _, err := tx.Exec(
		"INSERT INTO LoginTokens (user_id, content) VALUES ($1, $2);",
		userID,
		token,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) RemoveLoginToken(username string, token string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userID, err := getUserID(tx, username)
	if err != nil {
		return err
	}

	var exists bool
	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM LoginTokens WHERE user_id = $1 AND content = $2);",
		userID,
		token,
	).Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("no login token found for user '%s' with token '%s'", username, token)
	}

	if _, err := tx.Exec(
		"DELETE FROM LoginTokens WHERE user_id = $1 AND content = $2;",
		userID,
		token,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) HasLoginToken(username string, token string) (bool, error) {
	var exists bool

	tx, err := s.db.Begin()
	if err != nil {
		return false, nil
	}

	defer tx.Rollback()

	userID, err := getUserID(tx, username)
	if err != nil {
		return false, err
	}

	if err := tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM LoginTokens WHERE user_id = $1 AND content = $2);",
		userID,
		token,
	).Scan(&exists); err != nil {
		return false, err
	}

	if err := tx.Commit(); err != nil {
		return false, nil
	}

	return exists, nil
}

func (s *Store) RemoveAllLoginTokens(username string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userID, err := getUserID(tx, username)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		"DELETE FROM LoginTokens WHERE user_id = $1;",
		userID,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func getUserID(tx *sql.Tx, username string) (int, error) {
	var userID int
	if err := tx.QueryRow(
		"SELECT id FROM Users WHERE username = $1;",
		username,
	).Scan(&userID); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no user found with username '%s'", username)
		}
		return 0, err
	}

	return userID, nil
}

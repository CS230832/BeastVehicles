package admin

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

func (s *Store) AddAdmin(admin *types.AdminPayload) error {
	hashedPassword, err := auth.HashPassword(admin.Password)
	if err != nil {
		return err
	}

	_, err = s.db.Exec("INSERT INTO Admins (email, password, first_name, last_name, is_super) VALUES ($1, $2, $3, $4, $5);", admin.Email, hashedPassword, admin.FirstName, admin.LastName, admin.IsSuper)

	return err
}

func (s *Store) RemoveAdmin(email string) error {
	_, err := s.db.Exec("DELETE FROM Admins WHERE email = $1;", email)
	return err
}

func (s *Store) GetAdmin(email string) (*types.AdminPayload, error) {
	row := s.db.QueryRow("SELECT email, password, first_name, last_name, is_super FROM Admins WHERE email = $1;", email)

	if row == nil {
		return nil, fmt.Errorf("admin with email %s does not exist", email)
	}

	var admin = types.AdminPayload{}

	if err := row.Scan(&admin.Email, &admin.Password, &admin.FirstName, &admin.LastName, &admin.IsSuper); err != nil {
		return nil, err
	}

	return &admin, nil
}

func (s *Store) AddToken(email string, content string) error {
	_, err := s.db.Exec("INSERT INTO LoginTokens (email, content) VALUES ($1, $2);", email, content)

	return err
}

func (s *Store) RemoveToken(email string, content string) error {
	_, err := s.db.Exec("DELETE FROM LoginTokens WHERE email = $1 AND content = $2", email, content)
	return err
}

func (s *Store) GetTokens(email string) ([]string, error) {
	var tokens []string

	rows, err := s.db.Query("SELECT content FROM LoginTokens WHERE email = $1;", email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}
	
	return tokens, nil
}

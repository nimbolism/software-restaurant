package userdb

import (
	"database/sql"
	"fmt"
)

// User represents user data in the database
type User struct {
	UserID       string
	Username     string
	Email        string
	QRCode       string
	PhoneNumber  string
	NationalCode string
	// Add more fields as needed
}

// AddUser adds a new user to the database
func AddUser(user *User, db *sql.DB) error {
	// Prepare SQL statement
	stmt, err := db.Prepare("INSERT INTO users (user_id, username, email, qr_code, phone_number, national_code) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return fmt.Errorf("failed to prepare SQL statement for adding user: %v", err)
	}
	defer stmt.Close()

	// Execute SQL statement
	_, err = stmt.Exec(user.UserID, user.Username, user.Email, user.QRCode, user.PhoneNumber, user.NationalCode)
	if err != nil {
		return fmt.Errorf("failed to add user to the database: %v", err)
	}

	return nil
}

// GetUserByID retrieves user information from the database by user ID
func GetUserByID(userID string, db *sql.DB) (*User, error) {
	// Prepare SQL statement
	stmt, err := db.Prepare("SELECT user_id, username, email, qr_code, phone_number, national_code FROM users WHERE user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement for retrieving user information: %v", err)
	}
	defer stmt.Close()

	// Execute SQL statement
	var user User
	err = stmt.QueryRow(userID).Scan(&user.UserID, &user.Username, &user.Email, &user.QRCode, &user.PhoneNumber, &user.NationalCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user information from the database: %v", err)
	}

	return &user, nil
}

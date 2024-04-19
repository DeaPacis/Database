package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CreateUser(db *sql.DB) error {

	var email, login, password string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter full name:")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Println("Enter email:")
	fmt.Scanln(&email)

	fmt.Println("Enter login:")
	fmt.Scanln(&login)

	fmt.Println("Enter password:")
	fmt.Scanln(&password)

	query := `INSERT INTO Users (full_name, email, login, password) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, fullName, email, login, password)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	fmt.Println("User created successfully")

	return nil
}

func ReadUser(db *sql.DB) error {

	query := `SELECT * FROM Users`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read users: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Users':")
	fmt.Printf("%-10s | %-30s | %-30s | %-30s | %-30s\n", "ID", "Full Name", "Email", "Login", "Password")
	for rows.Next() {
		var userID int
		var fullName, email, login, password string
		if err := rows.Scan(&userID, &fullName, &email, &login, &password); err != nil {
			return fmt.Errorf("failed to read user row: %v", err)
		}
		fmt.Printf("%-10d | %-30s | %-30s | %-30s | %-30s\n", userID, fullName, email, login, password)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read users: %v", err)
	}

	return nil
}

func UpdateUser(db *sql.DB) error {

	var id int

	fmt.Println("Enter user ID to update:")
	fmt.Scanln(&id)

	query := `UPDATE Users SET full_name = $1, email = $2, login = $3, password = $4 WHERE user_id = $5`

	var email, login, password string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter full name:")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Println("Enter new email:")
	fmt.Scanln(&email)

	fmt.Println("Enter new login:")
	fmt.Scanln(&login)

	fmt.Println("Enter new password:")
	fmt.Scanln(&password)

	_, err := db.Exec(query, fullName, email, login, password, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	fmt.Println("User updated successfully")

	return nil
}

func DeleteUser(db *sql.DB) error {

	var id int

	fmt.Println("Enter user ID to delete:")
	fmt.Scanln(&id)

	query := `DELETE FROM Users WHERE user_id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	fmt.Println("User deleted successfully")
	return nil
}

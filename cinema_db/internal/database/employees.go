package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CreateEmployee(db *sql.DB) error {

	var position, phoneNumber string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter full name:")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Println("Enter employee position:")
	fmt.Scanln(&position)

	fmt.Println("Enter employee phone number:")
	fmt.Scanln(&phoneNumber)

	query := `INSERT INTO Employees (full_name, position, phone_number) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, fullName, position, phoneNumber)
	if err != nil {
		return fmt.Errorf("failed to create employee: %v", err)
	}

	fmt.Println("Employee created successfully")

	return nil
}

func ReadEmployee(db *sql.DB) error {

	query := `SELECT * FROM Employees`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read employees: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Employees':")
	fmt.Printf("%-10s | %-30s | %-20s | %-15s\n", "ID", "Full Name", "Position", "Phone Number")
	for rows.Next() {
		var employeeID int
		var fullName, position, phoneNumber string
		if err := rows.Scan(&employeeID, &fullName, &position, &phoneNumber); err != nil {
			return fmt.Errorf("failed to read employee row: %v", err)
		}
		fmt.Printf("%-10d | %-30s | %-20s | %-15s\n", employeeID, fullName, position, phoneNumber)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read employees: %v", err)
	}

	return nil
}

func UpdateEmployee(db *sql.DB) error {

	var employeeID int

	fmt.Println("Enter employee ID to update:")
	fmt.Scanln(&employeeID)

	var position, phoneNumber string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter full name:")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Println("Enter new employee position:")
	fmt.Scanln(&position)

	fmt.Println("Enter new employee phone number:")
	fmt.Scanln(&phoneNumber)

	query := `UPDATE Employees SET full_name = $1, position = $2, phone_number = $3 WHERE employee_id = $4`

	_, err := db.Exec(query, fullName, position, phoneNumber, employeeID)
	if err != nil {
		return fmt.Errorf("failed to update employee: %v", err)
	}

	fmt.Println("Employee updated successfully")

	return nil
}

func DeleteEmployee(db *sql.DB) error {

	var employeeID int

	fmt.Println("Enter employee ID to delete:")
	fmt.Scanln(&employeeID)

	query := `DELETE FROM Employees WHERE employee_id = $1`

	_, err := db.Exec(query, employeeID)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %v", err)
	}

	fmt.Println("Employee deleted successfully")

	return nil
}

package database

import (
	"database/sql"
	"fmt"
)

func CreateSession(db *sql.DB) error {

	var movieID, hallID, employeeID int
	var date, time string
	var ticketPrice float64

	fmt.Println("Enter movie ID:")
	fmt.Scanln(&movieID)

	fmt.Println("Enter hall ID:")
	fmt.Scanln(&hallID)

	fmt.Println("Enter employee ID:")
	fmt.Scanln(&employeeID)

	fmt.Println("Enter date (YYYY-MM-DD):")
	fmt.Scanln(&date)

	fmt.Println("Enter time (HH:MM):")
	fmt.Scanln(&time)

	fmt.Println("Enter ticket price:")
	fmt.Scanln(&ticketPrice)

	query := `INSERT INTO Sessions (movie_id, hall_id, employee_id, date, time, ticket_price) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query, movieID, hallID, employeeID, date, time, ticketPrice)
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	fmt.Println("Session created successfully")

	return nil
}

func ReadSession(db *sql.DB) error {

	query := `SELECT * FROM Sessions`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read sessions: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Sessions':")
	fmt.Printf("%-10s | %-10s | %-10s | %-10s | %-15s | %-10s | %-10s\n", "ID", "Movie ID", "Hall ID", "Employee ID", "Date", "Time", "Ticket Price")
	for rows.Next() {
		var sessionID, movieID, hallID, employeeID int
		var date, time string
		var ticketPrice float64
		if err := rows.Scan(&sessionID, &movieID, &hallID, &employeeID, &date, &time, &ticketPrice); err != nil {
			return fmt.Errorf("failed to read session row: %v", err)
		}
		fmt.Printf("%-10d | %-10d | %-10d | %-10d | %-15s | %-10s | %-10.2f\n", sessionID, movieID, hallID, employeeID, date, time, ticketPrice)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read sessions: %v", err)
	}

	return nil
}

func UpdateSession(db *sql.DB) error {

	var sessionID, movieID, hallID, employeeID int
	var date, time string
	var ticketPrice float64

	fmt.Println("Enter session ID to update:")
	fmt.Scanln(&sessionID)

	fmt.Println("Enter new movie ID:")
	fmt.Scanln(&movieID)

	fmt.Println("Enter new hall ID:")
	fmt.Scanln(&hallID)

	fmt.Println("Enter new employee ID:")
	fmt.Scanln(&employeeID)

	fmt.Println("Enter new date (YYYY-MM-DD):")
	fmt.Scanln(&date)

	fmt.Println("Enter new time (HH:MM):")
	fmt.Scanln(&time)

	fmt.Println("Enter new ticket price:")
	fmt.Scanln(&ticketPrice)

	query := `UPDATE Sessions SET movie_id = $1, hall_id = $2, employee_id = $3, date = $4, time = $5, ticket_price = $6 WHERE session_id = $7`

	_, err := db.Exec(query, movieID, hallID, employeeID, date, time, ticketPrice, sessionID)
	if err != nil {
		return fmt.Errorf("failed to update session: %v", err)
	}

	fmt.Println("Session updated successfully")

	return nil
}

func DeleteSession(db *sql.DB) error {

	var sessionID int

	fmt.Println("Enter session ID to delete:")
	fmt.Scanln(&sessionID)

	query := `DELETE FROM Sessions WHERE session_id = $1`

	_, err := db.Exec(query, sessionID)
	if err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}

	fmt.Println("Session deleted successfully")

	return nil
}

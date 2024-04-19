package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CreateReservation(db *sql.DB) error {

	var seatsNumber int
	var status string

	fmt.Println("Enter seats number:")
	fmt.Scanln(&seatsNumber)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter payment type:")
	paymentType, _ := reader.ReadString('\n')
	paymentType = strings.TrimSpace(paymentType)

	fmt.Println("Enter status:")
	fmt.Scanln(&status)

	query := `INSERT INTO Reservation (seats_number, payment_type, status) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, seatsNumber, paymentType, status)
	if err != nil {
		return fmt.Errorf("failed to create reservation: %v", err)
	}

	fmt.Println("Reservation created successfully")

	return nil
}

func ReadReservation(db *sql.DB) error {

	query := `SELECT * FROM Reservation`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read reservations: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Reservations':")
	fmt.Printf("%-10s | %-15s | %-15s | %-15s\n", "ID", "Seats Number", "Payment Type", "Status")
	for rows.Next() {
		var reservationID, seatsNumber int
		var paymentType, status string
		if err := rows.Scan(&reservationID, &seatsNumber, &paymentType, &status); err != nil {
			return fmt.Errorf("failed to read reservation row: %v", err)
		}
		fmt.Printf("%-10d | %-15d | %-15s | %-15s\n", reservationID, seatsNumber, paymentType, status)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read reservations: %v", err)
	}

	return nil
}

func UpdateReservation(db *sql.DB) error {

	var reservationID int

	fmt.Println("Enter reservation ID to update:")
	fmt.Scanln(&reservationID)

	var seatsNumber int
	var status string

	fmt.Println("Enter new seats number:")
	fmt.Scanln(&seatsNumber)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter payment type:")
	paymentType, _ := reader.ReadString('\n')
	paymentType = strings.TrimSpace(paymentType)

	fmt.Println("Enter new status:")
	fmt.Scanln(&status)

	query := `UPDATE Reservation SET seats_number = $1, payment_type = $2, status = $3 WHERE reservation_id = $4`

	_, err := db.Exec(query, seatsNumber, paymentType, status, reservationID)
	if err != nil {
		return fmt.Errorf("failed to update reservation: %v", err)
	}

	fmt.Println("Reservation updated successfully")

	return nil
}

func DeleteReservation(db *sql.DB) error {

	var reservationID int

	fmt.Println("Enter reservation ID to delete:")
	fmt.Scanln(&reservationID)

	query := `DELETE FROM Reservation WHERE reservation_id = $1`

	_, err := db.Exec(query, reservationID)
	if err != nil {
		return fmt.Errorf("failed to delete reservation: %v", err)
	}

	fmt.Println("Reservation deleted successfully")

	return nil
}

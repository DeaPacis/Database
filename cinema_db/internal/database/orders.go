package database

import (
	"database/sql"
	"fmt"
)

func CreateOrder(db *sql.DB) error {

	var userID, reservationID, sessionID int
	var totalPrice float64

	fmt.Println("Enter user ID:")
	fmt.Scanln(&userID)

	fmt.Println("Enter reservation ID:")
	fmt.Scanln(&reservationID)

	fmt.Println("Enter session ID:")
	fmt.Scanln(&sessionID)

	fmt.Println("Enter total price:")
	fmt.Scanln(&totalPrice)

	query := `INSERT INTO Orders (user_id, reservation_id, session_id, total_price) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, userID, reservationID, sessionID, totalPrice)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	fmt.Println("Order created successfully")

	return nil
}

func ReadOrder(db *sql.DB) error {

	query := `SELECT * FROM Orders`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read orders: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Orders':")
	fmt.Printf("%-10s | %-10s | %-15s | %-15s | %-15s\n", "ID", "User ID", "Reservation ID", "Session ID", "Total Price")
	for rows.Next() {
		var orderID, userID, reservationID, sessionID int
		var totalPrice float64
		if err := rows.Scan(&orderID, &userID, &reservationID, &sessionID, &totalPrice); err != nil {
			return fmt.Errorf("failed to read order row: %v", err)
		}
		fmt.Printf("%-10d | %-10d | %-15d | %-15d | %-15.2f\n", orderID, userID, reservationID, sessionID, totalPrice)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read orders: %v", err)
	}

	return nil
}

func UpdateOrder(db *sql.DB) error {

	var orderID int

	fmt.Println("Enter order ID to update:")
	fmt.Scanln(&orderID)

	var userID, reservationID, sessionID int
	var totalPrice float64

	fmt.Println("Enter new user ID:")
	fmt.Scanln(&userID)

	fmt.Println("Enter new reservation ID:")
	fmt.Scanln(&reservationID)

	fmt.Println("Enter new session ID:")
	fmt.Scanln(&sessionID)

	fmt.Println("Enter new total price:")
	fmt.Scanln(&totalPrice)

	query := `UPDATE Orders SET user_id = $1, reservation_id = $2, session_id = $3, total_price = $4 WHERE order_id = $5`

	_, err := db.Exec(query, userID, reservationID, sessionID, totalPrice, orderID)
	if err != nil {
		return fmt.Errorf("failed to update order: %v", err)
	}

	fmt.Println("Order updated successfully")

	return nil
}

func DeleteOrder(db *sql.DB) error {

	var orderID int

	fmt.Println("Enter order ID to delete:")
	fmt.Scanln(&orderID)

	query := `DELETE FROM Orders WHERE order_id = $1`

	_, err := db.Exec(query, orderID)
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)
	}

	fmt.Println("Order deleted successfully")

	return nil
}

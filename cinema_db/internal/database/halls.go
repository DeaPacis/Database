package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CreateHall(db *sql.DB) error {

	var hallType, screenSize string
	var capacity int

	fmt.Println("Enter hall type:")
	fmt.Scanln(&hallType)

	fmt.Println("Enter hall capacity:")
	fmt.Scanln(&capacity)

	fmt.Println("Enter screen size:")
	fmt.Scanln(&screenSize)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter audio system:")
	audioSystem, _ := reader.ReadString('\n')
	audioSystem = strings.TrimSpace(audioSystem)

	query := `INSERT INTO Halls (type, capacity, screen_size, audio_system) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, hallType, capacity, screenSize, audioSystem)
	if err != nil {
		return fmt.Errorf("failed to create hall: %v", err)
	}

	fmt.Println("Hall created successfully")

	return nil
}

func ReadHall(db *sql.DB) error {

	query := `SELECT * FROM Halls`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read halls: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Halls':")
	fmt.Printf("%-10s | %-15s | %-10s | %-15s | %-15s\n", "ID", "Type", "Capacity", "Screen Size", "Audio System")
	for rows.Next() {
		var hallID, capacity int
		var hallType, screenSize, audioSystem string
		if err := rows.Scan(&hallID, &hallType, &capacity, &screenSize, &audioSystem); err != nil {
			return fmt.Errorf("failed to read hall row: %v", err)
		}
		fmt.Printf("%-10d | %-15s | %-10d | %-15s | %-15s\n", hallID, hallType, capacity, screenSize, audioSystem)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read halls: %v", err)
	}

	return nil
}

func UpdateHall(db *sql.DB) error {

	var hallID int

	fmt.Println("Enter hall ID to update:")
	fmt.Scanln(&hallID)

	var hallType, screenSize string
	var capacity int

	fmt.Println("Enter new hall type:")
	fmt.Scanln(&hallType)

	fmt.Println("Enter new hall capacity:")
	fmt.Scanln(&capacity)

	fmt.Println("Enter new screen size:")
	fmt.Scanln(&screenSize)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter audio system:")
	audioSystem, _ := reader.ReadString('\n')
	audioSystem = strings.TrimSpace(audioSystem)

	query := `UPDATE Halls SET type = $1, capacity = $2, screen_size = $3, audio_system = $4 WHERE hall_id = $5`

	_, err := db.Exec(query, hallType, capacity, screenSize, audioSystem, hallID)
	if err != nil {
		return fmt.Errorf("failed to update hall: %v", err)
	}

	fmt.Println("Hall updated successfully")

	return nil
}

func DeleteHall(db *sql.DB) error {

	var hallID int

	fmt.Println("Enter hall ID to delete:")
	fmt.Scanln(&hallID)

	query := `DELETE FROM Halls WHERE hall_id = $1`

	_, err := db.Exec(query, hallID)
	if err != nil {
		return fmt.Errorf("failed to delete hall: %v", err)
	}

	fmt.Println("Hall deleted successfully")

	return nil
}

package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func CreateMovie(db *sql.DB) error {

	var genre string
	var releaseYear, duration int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter movie title:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Enter movie genre:")
	fmt.Scanln(&genre)

	fmt.Println("Enter release year:")
	fmt.Scanln(&releaseYear)

	fmt.Println("Enter movie duration:")
	fmt.Scanln(&duration)

	query := `INSERT INTO Movies (title, release_year, duration, genre) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, title, releaseYear, duration, genre)
	if err != nil {
		return fmt.Errorf("failed to create movie: %v", err)
	}

	fmt.Println("Movie created successfully")

	return nil
}

func ReadMovie(db *sql.DB) error {

	query := `SELECT * FROM Movies`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to read movies: %v", err)
	}
	defer rows.Close()

	fmt.Println("Table 'Movies':")
	fmt.Printf("%-10s | %-30s | %-15s | %-10s | %-20s\n", "ID", "Title", "Release Year", "Duration", "Genre")
	for rows.Next() {
		var movieID, releaseYear, duration int
		var title, genre string
		if err := rows.Scan(&movieID, &title, &releaseYear, &duration, &genre); err != nil {
			return fmt.Errorf("failed to read movie row: %v", err)
		}
		fmt.Printf("%-10d | %-30s | %-15d | %-10d | %-20s\n", movieID, title, releaseYear, duration, genre)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to read movies: %v", err)
	}

	return nil
}

func UpdateMovie(db *sql.DB) error {

	var movieID int

	fmt.Println("Enter movie ID to update:")
	fmt.Scanln(&movieID)

	var genre string
	var releaseYear, duration int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter movie title:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Enter new movie genre:")
	fmt.Scanln(&genre)

	fmt.Println("Enter new release year:")
	fmt.Scanln(&releaseYear)

	fmt.Println("Enter new movie duration:")
	fmt.Scanln(&duration)

	query := `UPDATE Movies SET title = $1, release_year = $2, duration = $3, genre = $4 WHERE movie_id = $5`

	_, err := db.Exec(query, title, releaseYear, duration, genre, movieID)
	if err != nil {
		return fmt.Errorf("failed to update movie: %v", err)
	}

	fmt.Println("Movie updated successfully")

	return nil
}

func DeleteMovie(db *sql.DB) error {

	var movieID int

	fmt.Println("Enter movie ID to delete:")
	fmt.Scanln(&movieID)

	query := `DELETE FROM Movies WHERE movie_id = $1`

	_, err := db.Exec(query, movieID)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %v", err)
	}

	fmt.Println("Movie deleted successfully")

	return nil
}

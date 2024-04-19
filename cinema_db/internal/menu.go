package internal

import (
	"cinema_db/internal/database"
	"database/sql"
	"fmt"
	"os"
)

func ChooseOption(db *sql.DB) {
	fmt.Println("Choose table:")
	fmt.Println("	E - Employees")
	fmt.Println("	H - Halls")
	fmt.Println("	M - Movies")
	fmt.Println("	O - Orders")
	fmt.Println("	R - Reservation")
	fmt.Println("	S - Sessions")
	fmt.Println("	U - Users")
	fmt.Println("Press q to quit")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "E":
		employeeOperation := ChooseOperation("Employees")
		switch employeeOperation {
		case 1:
			database.CreateEmployee(db)
		case 2:
			database.ReadEmployee(db)
		case 3:
			database.UpdateEmployee(db)
		case 4:
			database.DeleteEmployee(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "H":
		hallOperation := ChooseOperation("Halls")
		switch hallOperation {
		case 1:
			database.CreateHall(db)
		case 2:
			database.ReadHall(db)
		case 3:
			database.UpdateHall(db)
		case 4:
			database.DeleteHall(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "M":
		movieOperation := ChooseOperation("Movies")
		switch movieOperation {
		case 1:
			database.CreateMovie(db)
		case 2:
			database.ReadMovie(db)
		case 3:
			database.UpdateMovie(db)
		case 4:
			database.DeleteMovie(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "O":
		orderOperation := ChooseOperation("Orders")
		switch orderOperation {
		case 1:
			database.CreateOrder(db)
		case 2:
			database.ReadOrder(db)
		case 3:
			database.UpdateOrder(db)
		case 4:
			database.DeleteOrder(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "R":
		reservationOperation := ChooseOperation("Reservation")
		switch reservationOperation {
		case 1:
			database.CreateReservation(db)
		case 2:
			database.ReadReservation(db)
		case 3:
			database.UpdateReservation(db)
		case 4:
			database.DeleteReservation(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "S":
		sessionOperation := ChooseOperation("Sessions")
		switch sessionOperation {
		case 1:
			database.CreateSession(db)
		case 2:
			database.ReadSession(db)
		case 3:
			database.UpdateSession(db)
		case 4:
			database.DeleteSession(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "U":
		userOperation := ChooseOperation("Users")
		switch userOperation {
		case 1:
			database.CreateUser(db)
		case 2:
			database.ReadUser(db)
		case 3:
			database.UpdateUser(db)
		case 4:
			database.DeleteUser(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "q":
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

func ChooseOperation(tableName string) int {
	fmt.Printf("%s table menu:\n", tableName)
	fmt.Println("	1 - Create")
	fmt.Println("	2 - Read")
	fmt.Println("	3 - Update")
	fmt.Println("	4 - Delete")
	fmt.Println("Press 0 to quit")

	var choice int
	fmt.Scanln(&choice)

	return choice
}

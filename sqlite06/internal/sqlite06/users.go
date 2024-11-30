package sqlite06

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// get users list
// manipulate user data

var (
	Filename = ""
)

type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

func openConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", Filename)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func exists(username string) int {
	username = strings.ToLower(username)
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := -1

	stmt := fmt.Sprintf(`SELECT ID FROM Users WHERE Username = '%s'`, username)

	rows, _ := db.Query(stmt)
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("exists Scan()", err)
			return -1
		}

		userID = id
	}

	return userID
}

func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)

	db, err := openConnection()
	if err != nil {

		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", d.Username)
		return -1
	}

	stmt := `INSERT INTO Users VALUES (NULL, ?)`

	_, err = db.Exec(stmt, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	userID = exists(d.Username)
	if userID == -1 {
		return userID
	}

	stmt = `INSERT INTO Userdata VALUES(?, ?, ?, ?)`
	_, err = db.Exec(stmt, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("db Exec()", err)
		return -1
	}

	return userID
}

func DeleteUser(id int) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt := fmt.Sprintf(`SELECT Username FROM Users WHERE ID = %d`, id)

	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	var username string

	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return err
		}
	}

	if exists(username) != id {
		return fmt.Errorf("user with ID %d does not exist", id)
	}

	stmt = `DELETE FROM Userdata WHERE UserID = ?`
	_, err = db.Exec(stmt, id)
	if err != nil {
		return err
	}

	stmt = `DELETE from Users WHERE ID = ?`
	_, err = db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func ListUsers() ([]Userdata, error) {
	Data := []Userdata{}
	db, err := openConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT ID, Username, Name, Surname, Description FROM Users, Userdata WHERE Users.ID = Userdata.UserID`)
	if err != nil {
		return Data, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var name string
		var surname string
		var desc string

		err = rows.Scan(&id, &username, &name, &surname, &desc)
		temp := Userdata{ID: id, Username: username, Name: name, Surname: surname, Description: desc}

		Data = append(Data, temp)
		if err != nil {
			return nil, err
		}
	}

	return Data, nil
}

func UpdateUser(d Userdata) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID == -1 {
		return errors.New("user does not exist")
	}

	d.ID = userID
	stmt := `UPDATE Userdata SET Name = ?, Surname = ?, Description = ? WHERE UserID = ?`

	_, err = db.Exec(stmt, d.Name, d.Surname, d.Description, d.ID)
	if err != nil {
		return err
	}
	return nil
}

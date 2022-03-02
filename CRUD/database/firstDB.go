package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id    int
	Name  string
	Email string
	Role  string
}

func DBConnection(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to DB")
	}
	return db
}
func InsertRow(emp Employee, db *sql.DB) error {
	query := "INSERT INTO employee(name, email, role) VALUES (?, ?, ?)"
	_, err := db.Exec(query, emp.Name, emp.Email, emp.Role)
	if err != nil {
		return err
	}
	return nil
}

func Update(emp Employee, db *sql.DB) error {
	query := "UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return errors.New("prepare query not working")
	}
	defer stmt.Close()
	_, err = stmt.Exec(emp.Name, emp.Email, emp.Role, emp.Id)
	return err
}

func DeleteById(id int, db *sql.DB) error {
	query := "DELETE FROM employee WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return errors.New("prepare query not working")
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
func GetById(id int, db *sql.DB) (*Employee, error) {
	var emp Employee
	err := db.QueryRow("SELECT * from employee WHERE id = ?", id).Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Role)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

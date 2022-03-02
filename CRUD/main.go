package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	a "github.com/shivanithemathlete/CRUD/database"
)

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	db := a.DBConnection("empdb")
	defer db.Close()
	stmt, err := db.Prepare("create table if not exists employee(id int NOT NULL AUTO_INCREMENT, name varchar(50), email varchar(50), role varchar(20), PRIMARY KEY(id));")
	check(err)
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Table Created Successfully..")
	}

	// InsertRow
	empNew := a.Employee{Id: 1, Name: "Gama", Email: "gama@gmail.com", Role: "SDE-II"}
	err = a.InsertRow(empNew, db)
	check(err)
	// GetById
	emp, err := a.GetById(3, db)
	fmt.Println("Result of GetById:", emp)
	// Update
	fmt.Println("Updating Row")
	err = a.Update(a.Employee{Id: 1, Name: "uiefe", Email: "sdjwwdj@gmail.com", Role: "Intern"}, db)
	check(err)
	// Delete
	fmt.Println("Deleting Row")
	err = a.DeleteById(2, db)
	check(err)

}

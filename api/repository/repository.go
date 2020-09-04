package repository

import (
	"database/sql"
	"log"
)

type Employee struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
}

type EmployeeRepository struct {
	Connection *sql.DB
}

func (repo EmployeeRepository) List() ([]Employee, error) {
	employees := []Employee{}
	rows, err := repo.Connection.Query("SELECT FIRSTNME,LASTNAME,SEX FROM EMPLOYEE")
	if err != nil {
		log.Println("query error", err)
		return employees, err
	}
	for rows.Next() {
		var firstName, lastName, gender string
		rows.Scan(&firstName, &lastName, &gender)
		log.Println(firstName, "-", lastName, "-", gender)
		employees = append(employees, Employee{firstName, lastName, gender})
	}
	return employees, nil
}

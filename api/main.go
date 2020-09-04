package main

import (
	"database/sql"
	"db2/business"
	"db2/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {

	url := "protocol=tcpip;HOSTNAME=db;PORT=50000;UID=db2inst1;PWD=password;database=sample"
	db, err := sql.Open("go_ibm_db", url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	repository := repository.EmployeeRepository{
		Connection: db,
	}

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		employees, _ := repository.List()
		json.NewEncoder(w).Encode(employees)
	})

	http.HandleFunc("/employee/transform", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte(`{"error": "cannot parse data"}`))
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"firstname":  req.Firstname,
			"lastname":   req.Lastname,
			"gender":     req.Gender,
			"genderCode": business.GenderMapNumber(req.Gender),
		})
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}

type Request struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
}

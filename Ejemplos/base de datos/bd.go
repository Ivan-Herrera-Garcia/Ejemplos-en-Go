package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "LAPTOP-HRGDT8DN"
var port = 1433
var user = "sa"
var password = "1234"
var database = "ESTUDIANTES"

var db *sql.DB

func main() {
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	defer db.Close()
	selectversion()

	leerinfo()

}

func selectversion() {
	ctx := context.Background()

	err := db.PingContext(ctx)

	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed ", err.Error())
	}
	fmt.Printf("%s\n", result)

}

func leerinfo() (int, error) {
	ctx := context.Background()

	err := db.PingContext(ctx)

	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}
	tsql := fmt.Sprintf("select num_control, nombre, edad from INFORMACION;")

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal("Error reading rows: " + err.Error())
		return -1, err
	}

	defer rows.Close()

	var count int = 0

	for rows.Next() {
		var numcontrol, nombre string
		var edad int

		err := rows.Scan(&numcontrol, &nombre, &edad)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
			return -1, err
		}

		fmt.Printf("Numero control: %s, Nombre: %s, Edad: %d\n", numcontrol, nombre, edad)
		count++

	}
	return count, nil

}

package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/read", read)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}
func index(res http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(res, "Successfully Complete!")
	check(err)
}
func amigos(res http.ResponseWriter, req *http.Request) {
	var rows *sql.Rows
	rows, err = db.Query(`SELECT aName FROM amigos;`)
	check(err)
	defer rows.Close()

	var s, name string
	s = "RETRIEVED RECORDS \n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(res, s)
}
func create(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer(name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "Create table customer", n)
}

func insert(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO test01.customer VALUES("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "INSERTED RECORD", n)

}
func read(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM test01.customer;`)
	check(err)

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		fmt.Println(name)

		fmt.Fprintln(res, "RETRIEVED NAME", name)
	}
}

func update(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE test01.customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "UPDATE RECORD", n)
}

func delete(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM test01.customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "DELETE RECORD", n)
}
func drop(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE test01.customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(res, "DROPPED TABLE customer")
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

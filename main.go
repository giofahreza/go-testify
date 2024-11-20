// Example of using postgres in go language
package main

import (
	"database/sql"
	"fmt"

	c "github.com/giofahreza/go-calculate"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Addition 1 + 2 = ", c.Add(1, 2))
	fmt.Println("Subtraction 2 - 1 = ", c.Sub(2, 1))

	db_name := "postgres"
	db_user := "postgres.djnjkkjtfajdqbjeoekf"
	db_pass := "HNJBq9cpIHmp2AOL"
	db_host := "aws-0-ap-southeast-1.pooler.supabase.com"

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require", db_name, db_user, db_pass, db_host))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// get all table in public schema
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}

		fmt.Println(tableName)
	}

	// get all data from product table
	rows, err = db.Query("SELECT * FROM product")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int
		err = rows.Scan(&id, &name, &price)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, name, price)
	}
}

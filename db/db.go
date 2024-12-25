package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // ใช้ _ นำเข้าไลบรารีเพื่อให้ Go รู้จัก driver นี้
)

func main() {
	// เชื่อมต่อกับ SQLite แบบไม่ใช้ cgo
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// สร้างตารางและดำเนินการต่าง ๆ กับฐานข้อมูล
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users (name) VALUES ('Alice')")
	if err != nil {
		log.Fatal(err)
	}

	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name) // Output: Alice
}

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "modernc.org/sqlite" // ใช้ SQLite driver ที่ไม่ใช้ cgo
)

func main() {
	// เชื่อมต่อกับฐานข้อมูล SQLite (สามารถใช้ :memory: หรือไฟล์)
	db, err := sql.Open("sqlite", "db/cart.db") // กำหนดชื่อไฟล์ฐานข้อมูล
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	// อ่านคำสั่ง SQL จากไฟล์
	sqlBytes, err := ioutil.ReadFile("sql/order-go-comp.sql")
	if err != nil {
		log.Fatal(err)
	}

	// รันคำสั่ง SQL ในไฟล์
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized successfully!")
}

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	var (
		dbAddressFlag  string
		dbUsernameFlag string
		dbPasswordFlag string
		dbNameFlag     string
	)
	flag.StringVar(&dbAddressFlag, "dbAddress", "127.0.0.1:33104", "name and port of db")
	flag.StringVar(&dbUsernameFlag, "dbUsername", "root", "username to db")
	flag.StringVar(&dbPasswordFlag, "dbPassword", "example", "password for user to db")
	flag.StringVar(&dbNameFlag, "dbName", "example", "name to db to connect to")
	flag.Parse()
	cfg := mysql.Config{
		Addr:                 dbAddressFlag,
		User:                 dbUsernameFlag,
		Passwd:               dbPasswordFlag,
		DBName:               dbNameFlag,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening connection: %s", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("Error connecting to db: %s", err)
	}
	fmt.Println("Connected to db!")
	for {
		select {
		case <-time.Tick(time.Second * 3):
			_, err := conn.Exec("INSERT INTO test (test_data) VALUES (?)", "test123")
			if err != nil {
				log.Fatal("Error inserting record: %s", err)
			}
			fmt.Println("Record inserted!")
		}
	}

}

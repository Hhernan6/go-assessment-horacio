package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func main() {
	count := 0
	timeout := 10
	sleep := 5 * time.Second
	timeoutCounter := timeout
	dsn := os.Getenv("DISQO_MYSQL_USER") +
		":" +
		os.Getenv("DISQO_MYSQL_ROOT_PASSWORD") +
		"@tcp(" + "0.0.0.0" +
		":" + os.Getenv("DISQO_MYSQL_HOST_PORT") +
		")/" + os.Getenv("DISQO_MYSQL_DATABASE")
	log.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.(any))
	}

	for true {
		err = db.Ping()
		count++
		if err == nil {
			fmt.Println("connected!")
			_, err = db.Exec("CREATE TABLE IF NOT EXISTS badass_users (id BIGINT AUTO_INCREMENT PRIMARY KEY, first_name varchar(255), last_name varchar(255))")
			if err != nil {
				panic(err.(any))
			}
			log.Println("badass_users table is created successfully!")
			break
		}
		// 65 seconds
		if count == timeout {
			errMsg := fmt.Errorf("timing out our connection to the rds initialization")
			panic(errMsg.(any))
		}
		fmt.Println("waiting for rds to initialize....")
		fmt.Printf("\nattempts remaining: [%v]\n\n", timeoutCounter)
		time.Sleep(sleep)
		timeoutCounter--
	}
}

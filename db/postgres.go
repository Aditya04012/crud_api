package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *sql.DB{
	dsn:="postgres://postgres:xxxxxx@localhost:5432/goapp"
	db,err:=sql.Open("pgx",dsn)

	if err!=nil{
		log.Fatal(err)
	}

	err=db.Ping();
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("DB connected!!")
	return db
}
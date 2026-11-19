package main

import (
	"database/sql"
	"fmt"

	seed "github.com/Pradyumnasavadatti/poll/db"
	"github.com/Pradyumnasavadatti/poll/internals/env"
	"github.com/Pradyumnasavadatti/poll/internals/store"
	_ "github.com/lib/pq"
)

//this is used to run seeding
func main(){
	dbUrl := env.GetString("DB_ADDR","")
	db,err := sql.Open("postgres",dbUrl)
	if err!=nil{
		fmt.Println("Cant connect to db");
	}
	defer db.Close()

	store := store.NewStorage(db)

	err = seed.Seeding(db,store)

	if err!=nil{
		fmt.Printf("Something went wrong %s",err.Error())
	}

	fmt.Println("Successfully seeded the db")


}
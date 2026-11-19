package main

import (
	"fmt"

	"github.com/Pradyumnasavadatti/poll/internals/db"
	"github.com/Pradyumnasavadatti/poll/internals/env"
	"github.com/Pradyumnasavadatti/poll/internals/store"
)


func main(){
	config := &Config{
		Addr: env.GetString("PORT",":8091"),
		DBAddr: env.GetString("DB_ADDR","postgres://myuser:mypassword@localhost:5432/pol"),
	}

	db,err := db.New(config.DBAddr);
	if err!=nil{
		fmt.Println("Unable to connect to db")
	}
	defer db.Close()

	store := store.NewStorage(db)
	app:= &application{
		config: config,
		store: store,
	}

	mux := app.mux()
	if err := app.run(mux); err!=nil{
		fmt.Println("Error")
		fmt.Println(err.Error())
	}
}

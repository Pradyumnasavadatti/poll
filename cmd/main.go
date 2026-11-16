package main

import (
	"fmt"

	"github.com/Pradyumnasavadatti/poll/internals/db"
	"github.com/Pradyumnasavadatti/poll/internals/env"
)


func main(){
	config := &Config{
		Addr: env.GetString("PORT",":8091"),
		DBAddr: env.GetString("DB_ADDR","postgres://myuser:mypassword@localhost:5432/pol"),
	}
	app:= &application{
		config: config,
	}

	db,err := db.New(app.Addr);
	if err!=nil{
		fmt.Println("Unable to connect to db")
	}
	defer db.Close()

	mux := app.mux()
	if err := app.run(mux); err!=nil{
		fmt.Println("Error")
		fmt.Println(err.Error())
	}
}

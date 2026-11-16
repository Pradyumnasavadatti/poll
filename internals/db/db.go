package db

import "database/sql"

func New(addr string)(*sql.DB,error){
	db, err := sql.Open("postgres",addr)
	if err!=nil{
		return db,nil
	}
	return nil,err;
}
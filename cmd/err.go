package main

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error, errStr string){
	fmt.Printf("%s %s \n",errStr, err.Error())
	WriteJson(w,&Message{
		Message:errStr,
	})
}
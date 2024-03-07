package main

import (
	"context"
	"net/http"
	"time"

	// "encoding/json"
	"fmt"
	// "net/http"
	// "os"

	"github.com/jackc/pgx/v5"
	"github.com/rejath-chandran/backend-api/db"
	// "github.com/rejath-chandran/backend-api/db"
)
type Person struct {
    Name string
}
type Application struct{

}

func main(){

	mux:=http.NewServeMux()
	ctx,_ := context.WithTimeout(context.Background(),time.Millisecond*5000)
	conn, error:= pgx.Connect(ctx,"postgresql://admin:admin@165.232.179.121:5432/taskdb?sslmode=disable")
    queries := db.New(conn)
    if error!=nil{
		fmt.Println(error)
	}else{
		fmt.Println("connected",conn)
	}



	mux.HandleFunc("GET /api",func(w http.ResponseWriter, r *http.Request) {
		
		w.Write([]byte("rejath chandran"))
    })

	

	http.ListenAndServe(":8000",mux)
}
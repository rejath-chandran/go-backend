package main

import (
	"context"
	"fmt"
	"net/http"
)

func (app *Application)Home(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	fetchedAuthor, err := app.qry.AllUsers(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fetchedAuthor)
		
	w.Write([]byte("rejath chandran"))
}
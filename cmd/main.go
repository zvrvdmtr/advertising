package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zvrvdmtr/advertising/pkg/api"
	"github.com/zvrvdmtr/advertising/pkg/handler"
	"github.com/zvrvdmtr/advertising/pkg/models"
	"os"
)

func main() {
	conn, err := models.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Start migrating")
	_, err = conn.Exec(context.Background(), `create table ad (
		id serial primary key,
		name varchar(200) not null,
		description varchar(1000),
		price numeric(10, 2) not null,
		photos varchar(100)[],
		created timestamp not null
	)`)
	if err != nil {
		fmt.Printf("Migration error: " + err.Error())
	}
	fmt.Println("End migration")
	handler := handler.NewHandler()
	handler.HandleFunc("^/ad/[0-9]+$", api.GetAd(conn))
	handler.HandleFunc("/ads", api.GetList(conn))
	handler.HandleFunc("/create", api.CreateAd(conn))
	fmt.Println("Connection established on localhost:8000")
	http.ListenAndServe(":8000", handler)
}

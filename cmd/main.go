package main

import (
	"context"
	"fmt"
	"net/http"

	// "os"

	"github.com/zvrvdmtr/advertising/internal/handler"
	"github.com/zvrvdmtr/advertising/internal/repository"
	"github.com/zvrvdmtr/advertising/internal/services"
)

func main() {
	// conn, err := repository.InitDB(os.Getenv("DATABASE_URL"))
	conn, err := repository.InitDB("postgres://postgres:postgres@localhost:5432")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	// port := os.Getenv("PORT")
	fmt.Println("Start migrating")
	_, err = conn.Exec(context.Background(), `create table if not exists ad (
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
	pgsqlRepo := repository.NewPostgresAdRepository(conn)
	service := services.NewAdService(pgsqlRepo)
	newHandler := handler.NewAdHandler(service)
	fmt.Println("Connection established on localhost:8000")
	http.ListenAndServe(":8000", newHandler)
}

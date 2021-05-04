package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v4"
	"github.com/zvrvdmtr/advertising/pkg/api"
	"github.com/zvrvdmtr/advertising/pkg/handler"
)


func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/postgres")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	handler := handler.NewHandler()
	handler.HandleFunc("^/ad/[0-9]+$", api.GetAd(conn))
	handler.HandleFunc("/ads", api.GetList(conn))
	handler.HandleFunc("/create", api.CreateAd(conn))
	fmt.Println("Connection established on localhost:8000")
	http.ListenAndServe(":8000", handler)
}

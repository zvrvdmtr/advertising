package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/jackc/pgconn"

	"github.com/jackc/pgx/v4"
)

type Ad struct {
	Id          int64 		`json:"id"`
	Name        string		`json:"name"`
	Description string		`json:"description"`
	Price       float64		`json:"price"`
	Photos      []string	`json:"photos"`
	Created     time.Time	`json:"created"`
}

func (a *Ad) UnmarshalJSON(data []byte) error {
	validate := struct {
		Id          int64
		Name        *string
		Description string
		Price       *float64
		Photos      []string
		Created     time.Time
	}{}

	err := json.Unmarshal(data, &validate)
	if err != nil {
		return err
	}
	if validate.Name == nil {
		return fmt.Errorf("Required field Name is missing")
	}
	if validate.Price == nil {
		return fmt.Errorf("Required field Price is missing")
	}
	if len(validate.Photos) > 3 {
		return fmt.Errorf("Max 3 photo")
	}

	if err != nil {
		return err
	}

	a.Name = *validate.Name
	a.Description = validate.Description
	a.Price = *validate.Price
	a.Photos = validate.Photos
	a.Created = time.Now()
	return nil
}

type DbConnectionRow interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type DbConnection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

var conn pgx.Conn

func InitDB(databaseUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return conn, err
	}
	return conn, nil
}


func Get(conn DbConnectionRow, id int) (Ad, error) {
	var ad Ad

	row := conn.QueryRow(context.Background(), "SELECT * FROM ad where id = $1", id)
	err := row.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Photos, &ad.Created)
	if err != nil {
		return ad, err
	}

	return ad, nil
}

func All(conn DbConnection, pageNumber int) ([]Ad, error) {
	var ads []Ad
	rows, err := conn.Query(context.Background(), "SELECT * FROM ad LIMIT 10 OFFSET $1", pageNumber*10)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ad Ad
		err := rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Photos, &ad.Created)
		if err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	return ads, err
}

func Create(conn DbConnectionRow, ad Ad) (Ad, error) {
	var newAd Ad
	insertQuery := `INSERT INTO ad (name, description, price, photos, created) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	row := conn.QueryRow(context.Background(), insertQuery, ad.Name, ad.Description, ad.Price, ad.Photos, time.Now())
	err := row.Scan(&newAd.Id, &newAd.Name, &newAd.Description, &newAd.Price, &newAd.Photos, &newAd.Created)
	if err != nil {
		return newAd, err
	}
	return newAd, err
}

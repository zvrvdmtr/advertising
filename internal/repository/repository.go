package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/zvrvdmtr/advertising/internal/domain"
)

func InitDB(databaseUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

type PostgresAdRepository struct {
	Conn *pgx.Conn
}

func NewPostgresAdRepository(conn *pgx.Conn) domain.AdRepositoryIterface {
	return PostgresAdRepository{conn}
}

func (p PostgresAdRepository) Get(id int) (domain.Ad, error) {
	var ad domain.Ad

	row := p.Conn.QueryRow(context.Background(), "SELECT * FROM ad where id = $1", id)
	err := row.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Photos, &ad.Created)
	if err != nil {
		return ad, err
	}

	return ad, nil
}

func (p PostgresAdRepository) All(pageNumber int) ([]domain.Ad, error) {
	var ads []domain.Ad
	rows, err := p.Conn.Query(context.Background(), "SELECT * FROM ad LIMIT 10 OFFSET $1", pageNumber*10)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ad domain.Ad
		err := rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.Photos, &ad.Created)
		if err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	return ads, err
}

func (p PostgresAdRepository) Create(ad domain.Ad) (domain.Ad, error) {
	var newAd domain.Ad
	insertQuery := `INSERT INTO ad (name, description, price, photos, created) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	row := p.Conn.QueryRow(context.Background(), insertQuery, ad.Name, ad.Description, ad.Price, ad.Photos, time.Now())
	err := row.Scan(&newAd.Id, &newAd.Name, &newAd.Description, &newAd.Price, &newAd.Photos, &newAd.Created)
	if err != nil {
		return newAd, err
	}
	return newAd, err
}

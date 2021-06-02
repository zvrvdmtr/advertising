package domain

import (
	"encoding/json"
	"fmt"
	"time"
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

type AdRepositoryIterface interface {
	Get(id int) (Ad, error)
	All(pageNumber int) ([]Ad, error)
	Create(ad Ad) (Ad, error)
}

type AdServiceInterface interface {
	GetAds(pageNumber int) ([]Ad, error)
	GetAdById(id int, params []string) (Ad, error)
	CreateAd(ad Ad) (Ad, error)
}
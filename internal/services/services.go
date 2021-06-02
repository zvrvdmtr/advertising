package services

import (
	"time"
	"github.com/zvrvdmtr/advertising/internal/domain"
	"github.com/zvrvdmtr/advertising/internal/repository"
)

type AdDTO struct {
	Id          *int64 		`json:"id"`
	Name        *string		`json:"name"`
	Price       *float64	`json:"price"`
	Photos      *[]string	`json:"photos,omitempty"`
	Description *string		`json:"description,omitempty"`
	Created 	*time.Time	`json:"created,omitempty"`
}

type AdsDTO struct {
	Id 		int64
	Name 	string
	Price 	float64
	Photo 	string
}


func GetAds(conn repository.DbConnection, pageNumber int) ([]AdsDTO, error) {
	ads, err := repository.All(conn, pageNumber)
	var adsdto []AdsDTO
	for _, ad := range ads {
		dto := AdsDTO{Id: ad.Id, Name: ad.Name, Price: ad.Price}
		if len(ad.Photos) > 0 {
			dto.Photo = ad.Photos[0]
		}
		adsdto = append(adsdto, dto)
	}
	if err != nil {
		return adsdto, err
	}
	return adsdto, err
}

func GetAdById(conn repository.DbConnectionRow, id int, params []string) (AdDTO, error) {
	ad, err := repository.Get(conn, id)
	dto := AdDTO{Id: &ad.Id, Name: &ad.Name, Price: &ad.Price, Photos: nil, Created: &ad.Created}
	for _, param := range params {
		switch param {
		case "description":
			dto.Description = &ad.Description
		case "photos":
			dto.Photos = &ad.Photos
		}
	}
	if err != nil {
		return dto, err
	}
	return dto, nil
}

func CreateAd(conn repository.DbConnectionRow, ad domain.Ad) (domain.Ad, error) {
	newAd, err := repository.Create(conn, ad)
	if err != nil {
		return newAd, err
	}
	return newAd, nil
}
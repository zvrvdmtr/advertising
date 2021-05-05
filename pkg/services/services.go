package services

import (
	"time"

	"github.com/zvrvdmtr/advertising/pkg/models"
)

type AdDTO struct {
	Id          *int64 		`json:"id"`
	Name        *string		`json:"name"`
	Price       *float64	`json:"price"`
	Photos      *[]string	`json:"photos,omitempty"`
	Description *string		`json:"description,omitempty"`
	Created 	*time.Time	`json:"created,omitempty"`
}


func GetAds(conn models.DbConnection) ([]AdDTO, error) {
	var addtos []AdDTO
	ads, err := models.All(conn)
	for _, ad := range ads {
		addto := AdDTO{Id: &ad.Id, Name: &ad.Name, Price: &ad.Price}

		if len(ad.Photos) > 0 {
			temp := []string{ad.Photos[0]}
			addto.Photos = &temp
		}
		addtos = append(addtos, addto)
	}
	if err != nil {
		return addtos, err
	}
	return addtos, err
}

func GetAdById(conn models.DbConnection, id int, params []string) (AdDTO, error) {
	ad, err := models.Get(conn, id)
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

func CreateAd(conn models.DbConnection, ad models.Ad) (models.Ad, error) {
	newAd, err := models.Create(conn, ad)
	if err != nil {
		return newAd, err
	}
	return newAd, nil
}
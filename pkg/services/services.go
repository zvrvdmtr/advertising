package services

import (
	"github.com/zvrvdmtr/advertising/pkg/models"
)


func GetAds(conn models.DbConnection) ([]models.Ad, error) {
	ads, err := models.All(conn)
	if err != nil {
		return ads, err
	}
	return ads, err
}

func GetAdById(conn models.DbConnection, id int, params []string) (models.Ad, error) {
	ad, err := models.Get(conn, id, params)
	if err != nil {
		return ad, err
	}
	return ad, nil
}

func CreateAd(conn models.DbConnection, ad models.Ad) (models.Ad, error) {
	newAd, err := models.Create(conn, ad)
	if err != nil {
		return newAd, err
	}
	return newAd, nil
}
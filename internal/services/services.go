package services

import (
	"time"
	"github.com/zvrvdmtr/advertising/internal/domain"
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

type AdService struct {
	adRepo domain.AdRepositoryIterface
}

func NewAdService(repo domain.AdRepositoryIterface) domain.AdServiceInterface {
	return AdService{repo}
}

func (adService AdService) GetAds(pageNumber int) ([]domain.Ad, error) {
	ads, err := adService.adRepo.All(pageNumber)
	var adsdto []AdsDTO
	for _, ad := range ads {
		dto := AdsDTO{Id: ad.Id, Name: ad.Name, Price: ad.Price}
		if len(ad.Photos) > 0 {
			dto.Photo = ad.Photos[0]
		}
		adsdto = append(adsdto, dto)
	}
	// FIXME: Change return value
	if err != nil {
		return make([]domain.Ad, 0), err
	}
	return make([]domain.Ad, 0), err
}

func (adService AdService) GetAdById(id int, params []string) (domain.Ad, error) {
	ad, err := adService.adRepo.Get(id)
	dto := AdDTO{Id: &ad.Id, Name: &ad.Name, Price: &ad.Price, Photos: nil, Created: &ad.Created}
	for _, param := range params {
		switch param {
		case "description":
			dto.Description = &ad.Description
		case "photos":
			dto.Photos = &ad.Photos
		}
	}
	// FIXME: Change return value
	if err != nil {
		return domain.Ad{}, err
	}
	return domain.Ad{}, nil
}

func (adService AdService) CreateAd(ad domain.Ad) (domain.Ad, error) {
	newAd, err := adService.adRepo.Create(ad)
	if err != nil {
		return newAd, err
	}
	return newAd, nil
}

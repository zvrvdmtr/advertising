package services

import (
	"context"
	"testing"
	"reflect"
	"github.com/jackc/pgx/v4"
	"github.com/zvrvdmtr/advertising/internal/domain"
)

type MockRow struct {
	Id int
	Name string
	Description string
	Price float64
	Photos []string
}

var mockRow = MockRow{1, "Vaz 2101", "Hello!", 100.01, []string{"first", "second", "third"}}

func (mockRow MockRow) Scan(dest ...interface{}) error {
	id := dest[0].(*int64)
	name := dest[1].(*string)
	description := dest[2].(*string)
	price := dest[3].(*float64)
	photos := dest[4].(*[]string)
	*id = int64(mockRow.Id)
	*name = mockRow.Name
	*description = mockRow.Description
	*price = mockRow.Price
	*photos = mockRow.Photos
	return nil
}

type TestConnection struct{}

func (testConnection TestConnection) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return mockRow
}

func TestGetByIdWithParamDescription(test *testing.T) {

	conn := TestConnection{}
	addto, _ := GetAdById(conn, 1, []string{"description"})
	if (*addto.Description != mockRow.Description) {
		test.Errorf("got %v want %v", *addto.Description, mockRow.Description)
	}
}

func TestGetByIdWithParamPhotos(test *testing.T) {

	conn := TestConnection{}
	addto, _ := GetAdById(conn, 1, []string{"photos"})
	if !(reflect.DeepEqual(*addto.Photos, mockRow.Photos)) {
		test.Errorf("got %v want %v", *addto.Photos, mockRow.Photos)
	}
}

func TestGetByIdWithoutParam(test *testing.T) {

	conn := TestConnection{}
	addto, _ := GetAdById(conn, 1, make([]string, 0))
	if (addto.Photos != nil) || (addto.Description != nil){
		test.Errorf("got %v want %v", *addto.Photos, mockRow.Photos)
	}
}

func TestCreateAd(test *testing.T) {
	mockConnection := TestConnection{}
	newAd, _ := CreateAd(mockConnection, domain.Ad{})
	if (newAd.Id != int64(mockRow.Id)) || (newAd.Name != mockRow.Name) || (newAd.Price != mockRow.Price) {
		test.Errorf("got %v want %v", newAd, mockRow)
	}
}
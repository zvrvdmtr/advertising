package repository

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/zvrvdmtr/advertising/internal/domain"
)

type MockRow struct {
	Id int
	Name string
	Description string
	Price float64
}

func (mockRow MockRow) Scan(dest ...interface{}) error {
	id := dest[0].(*int64)
	name := dest[1].(*string)
	description := dest[2].(*string)
	price := dest[3].(*float64)
	*id = int64(mockRow.Id)
	*name = mockRow.Name
	*description = mockRow.Description
	*price = mockRow.Price
	return nil
}

type TestConnection struct{}

func (testConnection TestConnection) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	mock := MockRow{1, "Vaz 2101", "Hello!", 100.01}
	return mock
}

var expectedAd = domain.Ad{1, "Vaz 2101", "Hello!", 100.01, []string{"123"}, time.Now()}

func TestCreateAd(test *testing.T) {
	mockConnection := TestConnection{}
	mockAd := domain.Ad{}
	newAd, _ := Create(mockConnection, mockAd)
	if (newAd.Id != expectedAd.Id) || (newAd.Name != expectedAd.Name) || (newAd.Price != expectedAd.Price) {
		test.Errorf("got %v want %v", newAd, expectedAd)
	}
}

func TestGetAd(test *testing.T) {
	mockConnection := TestConnection{}
	ad, _ := Get(mockConnection, 1)
	if (ad.Id != expectedAd.Id) || (ad.Name != expectedAd.Name) || (ad.Price != expectedAd.Price) {
		test.Errorf("got %v want %v", ad, expectedAd)
	}
}
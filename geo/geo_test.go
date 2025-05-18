package geo_test

import (
	"practice/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange (Preparing conditions for the test)
	city := "London"
	expected := geo.GeoData{City: "London"}

	// Act (Action)
	got, err := geo.GetMyLocation(city)

	// Assert (Checking the result)
	if err != nil {
		t.Error(err)
	}

	if got.City != expected.City {
		t.Errorf("City expected %s got %s", expected.City, got.City)
	}
}

func TestGetLocationNoCity(t *testing.T) {
	city := "Londondon"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("City expected %s got %s", geo.ErrorNoCity, err)
	}
}

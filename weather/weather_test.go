package weather_test

import (
	"practice/weather/geo"
	"practice/weather/weather"
	"strings"
	"testing"
)

type testCasesStruct struct {
	name   string
	format int
}

func TestGetWeather(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{City: expected}
	format := 3

	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("City expected %s got %s", expected, result)
	}
}

var testCases = []testCasesStruct{
	{name: "Big format", format: 123},
	{name: "0 format", format: 0},
	{name: "negative format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{City: expected}

			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrorWrongFormat {
				t.Errorf("Expected %s got %s", weather.ErrorWrongFormat, err)
			}
		})
	}
}

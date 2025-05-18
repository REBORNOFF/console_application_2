package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

var ErrorNoCity = errors.New("NO_CITY")
var ErrorNot200 = errors.New("NOT200")

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrorNoCity
		}
		return &GeoData{City: city}, nil
	}

	response, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, ErrorNot200
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}

	return &geo, nil
}

func checkCity(city string) bool {
	postBody, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	urlData := "https://countriesnow.space/api/v0.1/countries/population/cities"
	response, err := http.Post(urlData, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var populationResponse CityPopulationResponse
	err = json.Unmarshal(body, &populationResponse)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return !populationResponse.Error
}

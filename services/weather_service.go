package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func weatherByCity(city string, country string, temprature bool, description bool, humidity bool) {
	var url strings.Builder
	url.WriteString("http://api.openweathermap.org/data/2.5/weather")
	url.WriteString("?q")
	url.WriteString(city)
	if len(country) > 0 {
		url.WriteString(",")
		url.WriteString(country)
	}

	response, err := http.Get(url.String())

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
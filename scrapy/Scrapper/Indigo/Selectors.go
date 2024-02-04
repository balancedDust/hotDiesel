package indigo

import (
	"io/ioutil"
	"strings"
)

var Selectors = map[string]string{
	"departureTime":    ".flight-departure-time > span",
	"arrivalTime":      ".flight-arrival-time > span",
	"durationTime":     "span.flight-duration__time",
	"flightStops":      "div.flight-stops",
	"flightNumbers":    "span.flight-number",
	"departureAirport": "span.flight-details__destinations__departure__airport",
	"arrivalAirport":   "span.flight-details__destinations__arrival__airport",
	"minimumFare":      "span.selected-fare__price",
	"availableFares":   "span.fare-type__amount__price__number"}

var FromCitySelector = "input[placeholder=From]"

var ToCitySelector = "input[placeholder=To]"

var FromDateSelector = "input[placeholder='Travel Dates']"

var SearchButtonSelector = "div.widget-container__search-form > button.custom-button"

var ResultsNotFoundSelector = "div.no-search-found__desc"

var ResultsSelector = "div.search-result-page__search-results__list__item"

var DefaultURL = "https://goindigo.com"

var CitiesMap = func() map[string]string {
	cityMap := map[string]string{}
	content, _ := ioutil.ReadFile("airports.txt")
	fileContent := strings.Split(string(content), ";")
	for _, city := range fileContent {
		cityWithNoCode := strings.Split(city, "(")[0]
		cityMap[cityWithNoCode] = city
	}
	return cityMap
}

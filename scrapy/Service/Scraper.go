package service

import (
	"fmt"

	dto "github.com/balancedDust/scrapy/Dto"
	scrapper "github.com/balancedDust/scrapy/Scrapper"
	indigo "github.com/balancedDust/scrapy/Scrapper/Indigo"
)

const defaultWait float64 = 300.0
const defaultSearchWait float64 = 1000.0

func inputParams(f scrapper.Scraper, input dto.Query) {
	f.Page.Locator("span#salePopupCloseBtn").Click()
	fromCity := f.Page.Locator(indigo.FromCitySelector)
	fromCity.Fill(input.FromCity)
	f.Page.WaitForTimeout(defaultWait)
	toCity := f.Page.Locator(indigo.ToCitySelector)
	toCity.Fill(input.ToCity)
	f.Page.WaitForTimeout(defaultWait)
	fromDate := f.Page.Locator(indigo.FromDateSelector)
	fromDate.Fill(input.FromDate)
	f.Page.WaitForTimeout(defaultWait)
	button := f.Page.Locator(indigo.SearchButtonSelector)
	button.Click()
}

func waitForResults(f scrapper.Scraper) error {
	f.Page.WaitForTimeout(defaultSearchWait)
	_, err := f.Page.Locator(indigo.ResultsNotFoundSelector).InnerText()
	if err == nil {
		return fmt.Errorf("no flights")
	} else {
		return nil
	}
}

func scrapeResults(f scrapper.Scraper) (resultsSlice []dto.FlightRoute) {
	results, _ := f.Page.Locator(indigo.ResultsSelector).All()
	for _, result := range results {
		depT, _ := result.Locator(indigo.Selectors["departureTime"]).InnerText()
		arrT, _ := result.Locator(indigo.Selectors["arrivalTime"]).InnerText()
		durT, _ := result.Locator(indigo.Selectors["durationTime"]).InnerText()
		fliS, _ := result.Locator(indigo.Selectors["flightStops"]).InnerText()
		minF, _ := result.Locator(indigo.Selectors["minimumFare"]).InnerText()
		depA, _ := result.Locator(indigo.Selectors["departureAirport"]).InnerText()
		arrA, _ := result.Locator(indigo.Selectors["arrivalAirport"]).InnerText()
		avlF, _ := result.Locator(indigo.Selectors["availableFares"]).AllInnerTexts()
		fliN, _ := result.Locator(indigo.Selectors["flightNumbers"]).AllInnerTexts()
		flightRoute := dto.FlightRouteBuilder().SetDepartureTime(depT).SetFlightStops(fliS).SetMinimumFare(minF).SetArrivalTime(arrT).SetDepartureAirport(depA).SetArrivalAirport(arrA).SetAvailableFares(avlF).SetFlightNumbers(fliN).SetDurationTime(durT)
		resultsSlice = append(resultsSlice, flightRoute)
	}
	return
}

func Scrape(input dto.Query) ([]dto.FlightRoute, error) {

	s := scrapper.Scraper{}
	s.Scraper().Browser(true).Goto(indigo.DefaultURL)
	s.Page.SetDefaultTimeout(6000)

	inputParams(s, input)

	err := waitForResults(s)
	if err != nil {
		fmt.Println("Error in Waiting for Results")
		return nil, err
	}

	results := scrapeResults(s)

	s.Finish()

	return results, nil
}

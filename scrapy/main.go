package main

import (
	"fmt"

	dto "github.com/balancedDust/scrapy/Dto"
	service "github.com/balancedDust/scrapy/Service"
)

func main() {
	query := dto.Query{QueryId: "643fg43f", RouteInput: dto.RouteInput{FromCity: "Mumbai(BOM)", ToCity: "Bengaluru(BLR)", FromDate: "7 Feb 2024"}}
	flights, err := service.Scrape(query)
	if err == nil {
		for _, flight := range flights {
			fmt.Print(flight)
		}
	}
}

package main

import (
	"fmt"

	dto "github.com/balancedDust/scrapy/dto"
	"github.com/balancedDust/scrapy/service"
)

func main() {
	queries := []dto.Query{dto.Query{QueryId: "643fg43f", RouteInput: dto.RouteInput{FromCity: "mumbai", ToCity: "bengaluru", FromDate: "7 Feb 2024"}},
		dto.Query{QueryId: "643f5443f", RouteInput: dto.RouteInput{FromCity: "jaypore", ToCity: "bangalore", FromDate: "1 Feb 2024"}}}
	valid, _ := service.ValidateInputs(queries)
	fmt.Println(service.Scrape(valid[0]))
}

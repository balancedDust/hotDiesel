package service

import (
	"strings"

	dto "github.com/balancedDust/scrapy/dto"
	indigo "github.com/balancedDust/scrapy/scrapper/indigo"
)

func ValidateInputs(i []dto.Query) (valid, invalid []dto.Query) {
	cities := indigo.CitiesMap()
	for _, input := range i {
		fC, tC := strings.Title(input.FromCity), strings.Title(input.ToCity)
		fCR, fPresent := cities[fC]
		tCR, tPresent := cities[tC]
		if !fPresent || !tPresent {
			invalid = append(invalid, input)
		} else {
			input.FromCity = fCR
			input.ToCity = tCR
			valid = append(valid, input)
		}
	}
	return valid, invalid
}

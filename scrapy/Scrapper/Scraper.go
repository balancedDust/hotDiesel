package scrapper

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

type Scraper struct {
	playwright *playwright.Playwright
	browser    playwright.Browser
	Page       playwright.Page
	Response   playwright.Response
}

func (s *Scraper) Scraper() *Scraper {
	if s.playwright != nil {
		log.Fatalf("scraper is set")
	}
	pw, err := playwright.Run()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	s.playwright = pw
	return s
}

func (s *Scraper) Browser(headless bool) *Scraper {
	if s.playwright == nil {
		log.Fatalf("scraper not set")
	}
	if s.browser != nil {
		log.Fatalf("browser is set")
	}
	options := playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(headless)}
	browser, err := s.playwright.Firefox.Launch(options)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	s.browser = browser
	return s
}

func (s *Scraper) Goto(url string) *Scraper {
	if s.browser == nil {
		log.Fatalf("browser not set")
	}
	if s.Page != nil {
		log.Fatalf("url is set")
	}
	page, err := s.browser.NewPage()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	s.Page = page
	response, err := s.Page.Goto(url)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	s.Response = response
	return s
}

func (s *Scraper) Finish() {
	s.Page.Close()
	s.browser.Close()
	s.playwright.Stop()
}

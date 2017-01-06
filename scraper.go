package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetCinemaList() []map[string]string {
	doc, err := goquery.NewDocument("http://www.unitedcinemas.com.au/session-times")
	if err != nil {
		log.Fatal(err)
	}

	var anchors *goquery.Selection
	var cinemas []map[string]string

	doc.Find("ul.nav.navbar-nav li.dropdown a").Each(func(i int, s *goquery.Selection) {
		if strings.TrimSpace(s.Text()) == "Session Times" {
			anchors = s.Parent().Find("ul.dropdown-menu li a")
		}
	})

	anchors.Each(func(i int, s *goquery.Selection) {
		cinema := map[string]string{}
		cinema["name"] = s.Text()
		link, _ := s.Attr("href")
		cinema["link"] = link
		parts := strings.Split(link, "/")
		cinema["id"] = parts[len(parts)-1]
		cinemas = append(cinemas, cinema)
	})

	return cinemas
}

func addAddress(cinema map[string]string) {
	doc, err := goquery.NewDocument(cinema["link"])
	if err != nil {
		log.Fatal(err)
	}

	body := []string{}
	well := doc.Find("div.well h3#session-details-title").First().Parent().Contents()
	well.Each(func(i int, s *goquery.Selection) {
		if s.Is("h3") {
			return
		}

		text := strings.TrimSpace(s.Text())
		if len(text) > 0 {
			body = append(body, text)
		}
	})

	cinema["address"] = strings.Join(body, ", ")
}

func main() {
	cinemas := GetCinemaList()

	for _, cinema := range cinemas {
		addAddress(cinema)
	}

	for _, cinema := range cinemas {
		fmt.Printf("%+v\n", cinema)
	}
}

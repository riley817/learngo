package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

var baseUrl string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=python&start="

type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageUrl := baseUrl + "&start=" + strconv.Itoa(page * 50)
	fmt.Println("Requesting : ", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCard := doc.Find(".jobsearch-SerpJobCard")

	searchCard.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := card.Find(".title>a").Text()
		location := card.Find(".sjcl").Text()
		fmt.Println(id, title, location)
	})


}

func getPages() int {
	pages := 0
	res, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length() - 1
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}

func cleanString(str string) string {
	
}

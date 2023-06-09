package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	p := "https://golang.org"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	doc, er := goquery.NewDocumentFromReader(re.Body)
	if er != nil {
		panic(er)
	}

	doc.Find("a").Each(func(i int, sel *goquery.Selection) {
		lk, _ := sel.Attr("href")
		println(i, sel.Text(), "(", lk, ")")
	})
}

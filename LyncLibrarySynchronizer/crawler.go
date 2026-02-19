package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
	"github.com/PuerkitoBio/goquery"
)

func fetchLCSC(url string) (string, string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	
	resp, err := client.Do(req)
	if err != nil { return "", "", err }
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	lcscID := ""
	re := regexp.MustCompile(`(\d+)\.html`)
	match := re.FindStringSubmatch(url)
	if len(match) > 1 { lcscID = "C" + match[1] }

	title := doc.Find("h1, .product-title").Text()
	params := ""
	doc.Find(".product-param-item, .param-table tr, .product-params tr, .attribute-item, table tr").Each(func(i int, s *goquery.Selection) {
		line := strings.TrimSpace(s.Text())
		if line != "" { params += line + " | " }
	})
	
	return fmt.Sprintf("Title: %s\nData: %s", title, params), lcscID, nil
}
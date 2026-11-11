package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func crawl(url string) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/157.0.0.0 Safari/537.36`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(bodyByte)
	fmt.Println("===============> ", url)

	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := a.Val
					if !strings.HasPrefix(link, "#") &&
						!strings.HasPrefix(link, "/") &&
						!strings.HasPrefix(link, "https") {
						crawl(url + "/" + link)
					}
					break
				}
			}
		}
	}
}

func main() {
	crawl("http://example.com")
}

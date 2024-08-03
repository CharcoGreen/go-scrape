package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// fetchHTML obtiene el contenido HTML desde la URL especificada
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: código de estado %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// parseHTML recorre el documento HTML y extrae los datos
func parseHTML(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				fmt.Println(attr.Val)
			}
		}
	}

	fmt.Println(n.Attr)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHTML(c)
	}
}

func main() {
	url := "https://elapuron.com"
	doc, err := fetchHTML(url)
	if err != nil {
		log.Fatalf("Error al obtener HTML: %v", err)
	}

	parseHTML(doc)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}
	url := os.Args[1]
	path := strings.Split(url, "/n/")
	id := path[len(path)-1]

	// HTTPリクエストを送信してHTMLを取得
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// 取得したHTMLをgoqueryでパース
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 記事のタイトルと内容を取得
	title := doc.Find(".o-noteContentHeader__title").Text()
	content := doc.Find(".note-common-styles__textnote-body").Text()

	// ファイルに出力
	fileContent, err := os.Create(id + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()

	_, err = fmt.Fprintln(fileContent, title)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(fileContent, content)
	if err != nil {
		log.Fatal(err)
	}
}

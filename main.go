package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"unicode/utf8"
)

type Param struct {
	source string
	target string
}

type Res struct {
	Status int    `json:"status"`
	Text   string `json:"text"`
}

func main() {
	api_url := os.Getenv("TRNS_URL")
	if api_url == "" {
		fmt.Println("You need to set api url on $TRANS_URL")
		return
	}
	flag.Parse()
	text := flag.Arg(0)
	if text == "" {
		fmt.Println("You need text to translate.")
		return
	}
	p := checkLang(text)

	url := fmt.Sprintf("%s?text=%s&source=%s&target=%s", api_url, url.QueryEscape(text), p.source, p.target)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %s", err.Error()))
		return
	}
	defer resp.Body.Close()

	var res Res
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %s", err.Error()))
		return
	}

	fmt.Println(res.Text)
}

func checkLang(text string) Param {
	if len(text) == utf8.RuneCountInString(text) {
		return Param{
			source: "en",
			target: "ja",
		}
	}
	return Param{
		source: "ja",
		target: "en",
	}
}

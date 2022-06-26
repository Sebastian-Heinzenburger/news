package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"
)

func print_regional(n int) string {
	resp_reg, err := http.Get("https://www.tagesschau.de/api2/news/?regions=1")

	if err != nil {
		return ""
	}
	defer resp_reg.Body.Close()

	body_reg, err := io.ReadAll(resp_reg.Body)
	if err != nil {
		return ""
	}

	var regional Regional

	json.Unmarshal(body_reg, &regional)

	var r strings.Builder
	r.WriteString("\033[4;1;95mRegionale Nachriten\033[24m\n")
	for i := 0; i < n; i++ {
		r.WriteString("\033[0;95m" + regional.News[i].Title + "\033[0m\n" + regional.News[i].FirstSentence + "\n")
	}

	return r.String()
}

func print_news(n int) string {
	resp_news, err := http.Get("https://www.tagesschau.de/api2/")

	if err != nil {
		return ""
	}
	defer resp_news.Body.Close()

	body_news, err := io.ReadAll(resp_news.Body)
	if err != nil {
		return ""
	}

	var news News

	json.Unmarshal(body_news, &news)

	var r strings.Builder
	r.WriteString("\033[4;1;36mAktuelle Nachriten\033[24m\n")

	l := 0
	for i := 0; i < n; i++ {
		if l < utf8.RuneCountInString(news.News[i].Topline) {
			l = utf8.RuneCountInString(news.News[i].Topline)
		}
	}
	for i := 0; i < n; i++ {
		str := "\033[0;36m" + news.News[i].Topline + strings.Repeat(" ", l-utf8.RuneCountInString(news.News[i].Topline)+1) + "\033[0m" + news.News[i].FirstSentence + "\n"
		r.WriteString(str)
	}

	return r.String()
}

func main() {

	var flag_regional bool
	var flag_all bool
	var n int

	flag.BoolVar(&flag_regional, "r", false, "Print Regional News for Baden-Württemberg")
	flag.BoolVar(&flag_all, "a", false, "Print Regular and Regional News")
	flag.IntVar(&n, "n", 10, "Number of articles to show")
	flag.Parse()

	if flag_regional {
		fmt.Println(print_regional(n))
		return
	} else if flag_all {
		fmt.Println(print_news(n))
		fmt.Println(print_regional(n))
	} else {
		fmt.Println(print_news(n))
	}

}

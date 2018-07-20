package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Articles struct {
	Articles []Article `json:"articles"`
}

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/tass", loadTASSFeed)
	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Printf("Listening at http://localhost:%s ...", PORT)
	}

}

func loadTASSFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		url := ("http://tass.ru/api/news/lenta?limit=50&sections[]=4964&not_save=1")

		var a Articles
		if err := getJSON(url, &a); err != nil {
			fmt.Println(err)
			return
		}

		t, _ := template.ParseFiles("feed.gtpl")
		t.Execute(w, a.Articles)
	} else {
		fmt.Fprintf(w, "Pease access page using GET method")
	}

}

func getJSON(url string, articles *Articles) error {
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if err = json.Unmarshal([]byte(buf), &articles); err != nil {
			return err
		}
		for i, j := range articles.Articles {
			articles.Articles[i].Url = "http://tass.ru" + j.Url
		}
	} else {
		return fmt.Errorf("HTTP response status ode: %d\tMessage: %s", res.StatusCode, res.Status)
	}
	return nil
}

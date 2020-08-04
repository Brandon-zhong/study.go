package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {

	testJson()

}

func testDecoder() {
	result, err := SearchIssues(strings.Split("repo:golang/go is:open json decoder", " "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func testJson() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON marshaling failed: %s", err)
	}
	fmt.Println(string(data))

	jsonStr := "[{\"Title\":\"java\",\"released\":2010,\"Actors\":null},{\"Title\":\"python\",\"released\":2015,\"color\":true,\"Actors\":null},{\"Title\":\"go\",\"released\":2020,\"Actors\":null}]"

	fmt.Println("--------")

	var m []Movie
	_ = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&m)
	fmt.Println(m)
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "java", Year: 2010, Color: false},
	{Title: "python", Year: 2015, Color: true},
	{Title: "go", Year: 2020, Color: false},
}

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/parnurzeal/gorequest"
)

var api string
var token string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	api = os.Getenv("GITHUB_API")
	token = os.Getenv("GITHUB_TOKEN")
}

func fetchReactions(owner, repo string, issueNumber int, content string) []string {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d/reactions", api, owner, repo, issueNumber)

	request := gorequest.New()
	_, body, errs := request.Get(url).
		Set("Accept", "application/vnd.github.squirrel-girl-preview+json").
		Set("Authorization", fmt.Sprintf("token %s", token)).
		Param("content", content).
		End()
	if errs != nil {
		panic(errs)
	}

	var reactions []map[string]interface{}
	json.Unmarshal([]byte(body), &reactions)

	var usernames []string
	for _, reaction := range reactions {
		user, ok := reaction["user"].(map[string]interface{})
		if ok {
			login := user["login"]
			usernames = append(usernames, login.(string))
		}
	}

	return usernames
}

func fetchUserInfo(ch chan<- string, username string) {
	url := fmt.Sprintf("%s/users/%s", api, username)

	request := gorequest.New()
	_, body, errs := request.Get(url).
		Set("Accept", "application/vnd.github.squirrel-girl-preview+json").
		Set("Authorization", fmt.Sprintf("token %s", token)).
		End()
	if errs != nil {
		panic(errs)
	}

	var userInfo map[string]interface{}
	json.Unmarshal([]byte(body), &userInfo)

	ch <- userInfo["name"].(string)
}

func printName(ch <-chan string) {
	name := <-ch
	fmt.Println(name)
}

func main() {
	owner := flag.String("owner", "", "repository owner")
	repo := flag.String("repo", "", "repository name")
	issueNumber := flag.Int("issueNumber", 0, "issue number")
	content := flag.String("content", "", "content [+1, -1, heart, laugh, confused, hooray]")

	flag.Parse()

	ch := make(chan string, 4)

	usernames := fetchReactions(*owner, *repo, *issueNumber, *content)
	for _, username := range usernames {
		go fetchUserInfo(ch, username)
	}

	for _ = range usernames {
		fmt.Println(<-ch)
	}

	fmt.Println("Total:", len(usernames))
}

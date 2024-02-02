package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type github struct {
	Name     string `json:"login"`
	NumRepos uint16 `json:"public_repos"`
}

func gitHubInfo(user string) (string, uint16, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/" + user))
	if err != nil {
		return "", 0, fmt.Errorf("couldn't reach the server: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("couldn't fetch user: %w", err)
	}
	var info github
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&info); err != nil {
		return "", 0, fmt.Errorf("couldn't decode the response: %w", err)
	}
	return info.Name, uint16(info.NumRepos), err
}

func main() {
	user := flag.String("user", "", " please provide github user")
	flag.Parse()
	login := url.PathEscape(*user)
	name, repos, err := gitHubInfo(login)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	log.Println(name, repos)
}

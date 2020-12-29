package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api.github.com/user/repos"
	method := "POST"

	payload := strings.NewReader(`{

    "name" : "CreateRepoUsingGithubAPI" ,
    "description" : "Displays todays weather based on the specified location" ,
    "homepage" : "www.github.com"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Generate your Personal Access token on github using https://github.com/settings/tokens
	// Add your-Personal-access-token-generated-on-github
	req.Header.Add("Authorization", "Bearer your-Personal-access-token-generated-on-github")

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

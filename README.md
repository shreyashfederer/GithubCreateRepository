# GithubCreateRepository

Creates a Repository on Github from Golang Code using Github API - https://api.github.com/user/repos

1. Generate your Personal Access token on github -  https://github.com/settings/tokens

2. Add the Bearer token in request header using ``` req.Header.Add("Authorization", "Bearer your-Personal-access-token-generated-on-github")  ``` 

3. Make a POST Request to the endpoint -  https://api.github.com/user/repos .

    ### Sample Payload for Post Request

    ```
    {

    "name" : "CreateRepoUsingGithubAPI" ,
    "description" : "This Repository is created using Github API" ,
    "homepage" : "www.github.com"

    }

    ```

    ### Sample CUrl Request Format

    ```
    curl --location --request POST 'https://api.github.com/user/repos' \
    --header 'Accept: application/vnd.github.v3+json' \
    --header 'Authorization: Bearer yourbearerToken' \
    --header 'Content-Type: application/json' \
    --data-raw '{

    "name" : "CreateRepoUsingGithubAPI" ,
    "description" : "This Repository is created using Github API" ,
    "homepage" : "www.github.com"
}'

    ```

Usage:  ``` go run createrepository.go ```

To contribute:

* Fork the repo.
* Make changes in a new branch.
* Create a Pull Request.

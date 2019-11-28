package main

import (
	"context"
	"sort"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/static"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v28/github"

	"fmt"

	model "go-blog/models"

	"strconv"

	"github.com/gin-contrib/cors"
)

const github_token = "f46782677e93156562f72987229b85b5a82fef94"

func initGithubClient(accessToken string) (context.Context, *github.Client) {
	fmt.Printf("Initilizing Github Client")

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return ctx, client
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:5000"}

	r.Use(cors.New(config))

	ctx, github := initGithubClient(github_token)

	r.Use(static.Serve("/", static.LocalFile("client/public", true)))

	r.GET("/user", func(c *gin.Context) {
		user, _, _ := github.Users.Get(ctx, "akshay5995")

		c.JSON(200, user)
	})

	r.GET("/repos", func(c *gin.Context) {
		repos, _, _ := github.Repositories.List(ctx, "", nil)

		queryCount := c.Request.URL.Query().Get("count")

		var count int = 5

		var IsLoadedAll = false

		if queryCount != "" {
			count, _ = strconv.Atoi(queryCount)
		}

		var myRepos []model.Repo

		for _, val := range repos {
			if *val.Private || *val.Fork {
				continue
			}
			var r model.Repo
			r.Name = val.Name
			r.Fullname = val.FullName
			r.Url = val.HTMLURL
			r.Stars = val.StargazersCount
			r.ForksCount = val.ForksCount

			myRepos = append(myRepos, r)
		}

		sort.Slice(myRepos[:], func(i, j int) bool {
			return *myRepos[i].Stars > *myRepos[j].Stars
		})

		if count > len(myRepos) {
			count = len(myRepos)
			IsLoadedAll = true
		}

		c.JSON(200, gin.H{
			"repos":     myRepos[:count],
			"loadedAll": IsLoadedAll,
		})

	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":4001")
}

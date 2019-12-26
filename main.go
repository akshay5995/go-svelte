package main

import (
	"context"

	"net/http"

	"log"

	"sort"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"

	"github.com/gin-contrib/static"

	"golang.org/x/oauth2"

	"github.com/gin-contrib/logger"

	"github.com/google/go-github/v28/github"

	"fmt"

	model "go-svelte/models"

	"strconv"

	"os"

	"github.com/gin-contrib/cors"
)

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
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().
		Logger()

	_, connectionErr := http.Get("https://api.github.com")

	if connectionErr != nil {
		log.Fatal("Error connecting to github")
	}

	var githubToken = os.Getenv("GITHUB_TOKEN")

	var githubUser = os.Getenv("GITHUB_USER")

	r := gin.Default()

	r.LoadHTMLGlob("public/templates/*")

	r.Use(logger.SetLogger(logger.Config{
		Logger: &subLog,
		UTC:    true,
	}))

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:5000"}

	r.Use(cors.New(config))

	ctx, github := initGithubClient(githubToken)

	// Route for JS assets
	r.Use(static.Serve("/build", static.LocalFile("public/build", true)))

	// Route for Images
	r.Use(static.Serve("/img", static.LocalFile("public/img", true)))

	// Route for index.html and config object
	r.GET("/", func(c *gin.Context) {
		fmt.Printf("/ route called")

		var siteName = os.Getenv("SITE_NAME")

		var devBlogSite = os.Getenv("DEV_BLOG_SITE")

		//render with master
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": siteName,
			"blog":  devBlogSite,
			"user":  githubUser,
		})
	})

	r.GET("/user", func(c *gin.Context) {
		user, _, _ := github.Users.Get(ctx, githubUser)

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

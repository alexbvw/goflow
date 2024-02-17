package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v50/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func GithubAuthentication(c *fiber.Ctx) *github.Client {
	godotenv.Load()
	ctx := c.Context()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

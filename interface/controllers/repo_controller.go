package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v50/github"
)

func GetRepos(c *fiber.Ctx) error {
	ctx := c.Context()
	client := GithubAuthentication(c)
	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(&repos)
}

func GetGitUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	client := GithubAuthentication(c)
	// list all users for the authenticated user
	users, _, err := client.Users.ListAll(ctx, &github.UserListOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(&users)
}

func GetGitEvent(c *fiber.Ctx) error {
	ctx := c.Context()
	client := GithubAuthentication(c)
	// list all activities for the authenticated user
	events, _, err := client.Activity.ListEvents(ctx, &github.ListOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(&events)
}

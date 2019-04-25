package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var GITHUB_OWNER = os.Getenv("GITHUB_OWNER")

var app = cli.NewApp()

func commands() {
	app.Commands = []cli.Command{
		{
			Name:   "branch",
			Action: listBranches,
		},
		{
			Name:   "pbranch",
			Action: listPullRequestBranches,
		},
	}
}

func newClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GITHUB_TOKEN},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func listBranches(c *cli.Context) error {
	client := newClient()

	opt := &github.ListOptions{}
	branches, _, err := client.Repositories.ListBranches(context.Background(), GITHUB_OWNER, c.Args()[0], opt)
	if err != nil {
		return err
	}

	for _, b := range branches {
		fmt.Println(*b.Name)
	}

	return nil
}

func listPullRequestBranches(c *cli.Context) error {
	client := newClient()

	opt := &github.PullRequestListOptions{}
	pullrequests, _, err := client.PullRequests.List(context.Background(), GITHUB_OWNER, c.Args()[0], opt)
	if err != nil {
		return err
	}

	for _, p := range pullrequests {
		fmt.Println(p.GetHead().GetRef())
	}

	return nil
}

func main() {
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

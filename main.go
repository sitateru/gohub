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

var GITHUB_TOKEN = ""
var GITHUB_OWNER = ""

func main() {
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	GITHUB_OWNER = os.Getenv("GITHUB_OWNER")

	app := cli.NewApp()
	app.Action = listBranches

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func listBranches(c *cli.Context) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GITHUB_TOKEN},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.ListOptions{}
	branches, _, err := client.Repositories.ListBranches(ctx, GITHUB_OWNER, c.Args()[0], opt)
	if err != nil {
		return err
	}

	for _, b := range branches {
		fmt.Println(*b.Name)
	}

	return nil
}

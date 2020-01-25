package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/tjarratt/babble"
	"github.com/urfave/cli/v2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func main() {
	app := &cli.App{
		Name: "git-generate",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "commits",
				Aliases: []string{"c"},
				Value:   1,
				Usage:   "generate number of commits",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "test.txt",
				Usage:   "name of the generated file",
			},
			&cli.StringFlag{
				Name:  "prefix",
				Value: "",
				Usage: "prefix for commit messages",
			},
		},
		Action: func(c *cli.Context) error {
			r, err := git.PlainOpen(".")
			if err != nil {
				return err
			}

			w, err := r.Worktree()
			if err != nil {
				return err
			}

			f, err := os.OpenFile(c.String("file"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			defer f.Close()

			babbler := babble.NewBabbler()

			for i := 0; i < c.Int("commits"); i++ {
				if _, err := f.WriteString(fmt.Sprintf("%s\n", babbler.Babble())); err != nil {
					return err
				}

				if _, err := w.Add(c.String("file")); err != nil {
					return err
				}

				commitMsg, err := generateCommitMsg()
				if err != nil {
					return err
				}

				commit, err := w.Commit(fmt.Sprintf("%s%s", c.String("prefix"), commitMsg), &git.CommitOptions{
					Author: &object.Signature{
						Name:  "John Doe",
						Email: "john@doe.org",
						When:  time.Now(),
					},
				})

				if _, err := r.CommitObject(commit); err != nil {
					return err
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("fatal error: %s", err)
	}
}

func generateCommitMsg() (string, error) {
	res, err := http.Get("http://whatthecommit.com/index.txt")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

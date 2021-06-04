package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"example.com/m/unit4/issue"
)

var (
	create = flag.Bool("c", false, "")
	list   = flag.Bool("l", false, "")
	read   = flag.Bool("r", false, "")
	edit   = flag.Bool("e", false, "")

	owner  = flag.String("owner", "", "")
	repo   = flag.String("repo", "", "")
	number = flag.String("number", "", "")
	token  = flag.String("token", "", "")

	title = flag.String("title", "", "")
	body  = flag.String("body", "", "")
)

func main4_11() {
	flag.Parse()
	switch {
	case *list:
		p := issue.Params{Owner: *owner, Repo: *repo}
		issues, err := p.GetIssues()
		if err != nil {
			log.Fatal(err)
		}
		for _, issue := range issues {
			fmt.Printf("%s\t%s\n", issue.Title, issue.Body)
		}
	case *read:
		p := issue.Params{Owner: *owner, Repo: *repo, Number: *number}
		issue, err := p.GetIssue()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\n", issue.Title, issue.Body)
	case *create:
		p := issue.Params{Owner: *owner, Repo: *repo, Token: *token,
			Issue: issue.Issue{Title: *title, Body: *body}}
		if !p.CreateIssue() {
			fmt.Fprintf(os.Stderr, "create failed\n")
		}
	case *edit:
		p := issue.Params{Owner: *owner, Repo: *repo,
			Token: *token, Number: *number,
			Issue: issue.Issue{Title: *title, Body: *body}}
		if !p.EditIssue() {
			fmt.Fprintf(os.Stderr, "edit issue fail")
		}
	}
}

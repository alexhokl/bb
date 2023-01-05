package models

import (
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/datetime"
	"github.com/fatih/color"
)

// PrintID prints ID of a pull request
func PrintID(pr *swagger.Pullrequest) {
	fmt.Printf("%d\n", pr.Id)
}

// PrintShortDescription prints a short description of the specified pull request
func PrintShortDescription(pr *swagger.Pullrequest, isIncludeCreatedOn bool) {
	PrintOneLiner(pr)

	if isIncludeCreatedOn {
		color.New(color.FgGreen).Printf("\tCreated:%s\n", datetime.GetLocalDateTimeString(&pr.UpdatedOn))
	}

	color.New(color.FgMagenta).Printf("\t%s",
		pr.Source.Branch.Name)
	fmt.Printf(" -> ")
	color.New(color.FgYellow).Printf("%s\n",
		pr.Destination.Branch.Name)

	approveStr := ""
	for _, reviewer := range pr.Participants {
		if reviewer.Approved {
			approveStr = fmt.Sprintf("%s\n\tApproved by %s", approveStr, reviewer.User.DisplayName)
		}
	}
	if approveStr != "" {
		color.New(color.FgBlue).Println(approveStr)
	}
	fmt.Println("---")
}

// PrintOneLiner prints a short description of the specified pull request
func PrintOneLiner(pr *swagger.Pullrequest) {
	fmt.Printf("%d", pr.Id)
	color.New(color.FgGreen).Printf(" Updated:%s", datetime.GetLocalDateTimeString(&pr.UpdatedOn))
	color.New(color.FgCyan).Printf(" %s", pr.Author.DisplayName)
	fmt.Printf(" %s\n", pr.Title)
}

// PrintDescription prints description of the specified pull request
func PrintDescription(pr *swagger.Pullrequest) {
	// Todo - print description
	// fmt.Printf("Description:\n%s\n", pr.Description)
}

// PrintID prints ID of a pull request
func (pr PullRequestDetail) PrintID() {
	fmt.Printf("%d\n", pr.ID)
}

// PrintShortDescription prints a short description of the specified pull request
func (pr PullRequestDetail) PrintShortDescription(isIncludeCreatedOn bool) {
	pr.PrintOneLiner()

	if isIncludeCreatedOn {
		color.New(color.FgGreen).Printf("\tCreated:%s\n", datetime.GetLocalDateTimeString(&pr.UpdatedOn))
	}

	color.New(color.FgMagenta).Printf("\t%s",
		pr.Source.Branch.Name)
	fmt.Printf(" -> ")
	color.New(color.FgYellow).Printf("%s\n",
		pr.Destination.Branch.Name)

	approveStr := ""
	for _, reviewer := range pr.Participants {
		if reviewer.Approved {
			approveStr = fmt.Sprintf("%s\n\tApproved by %s", approveStr, reviewer.User.DisplayName)
		}
	}
	if approveStr != "" {
		color.New(color.FgBlue).Println(approveStr)
	}
	fmt.Println("---")
}

// PrintOneLiner prints a short description of the specified pull request
func (pr PullRequestDetail) PrintOneLiner() {
	fmt.Printf("%d", pr.ID)
	color.New(color.FgGreen).Printf(" Updated:%s", datetime.GetLocalDateTimeString(&pr.UpdatedOn))
	color.New(color.FgCyan).Printf(" %s", pr.Author.DisplayName)
	fmt.Printf(" %s\n", pr.Title)
}

// PrintDescription prints description of the specified pull request
func (pr PullRequestDetail) PrintDescription() {
	fmt.Printf("Description:\n%s\n", pr.Description)
}

// ToString returns the description of a comment
func (c Comment) ToString() string {
	return fmt.Sprintf(
		"Comment by %s (%s): %s",
		c.User.DisplayName,
		datetime.GetLocalDateTimeString(&c.UpdatedOn),
		c.Content.Raw)
}

// ToString returns the description of an update
func (c Update) ToString() string {
	return fmt.Sprintf(
		"Commit made by %s (%s): %s",
		c.Author.DisplayName,
		datetime.GetLocalDateTimeString(&c.Date),
		c.Source.Commit.Hash)
}

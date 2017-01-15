package models

import (
	"fmt"
	"time"
)

// PullRequest interface
type PullRequest interface {
	ToString() string
}

// PullRequestList struct
type PullRequestList struct {
	PageLen int               `json:"pagelen"`
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Items   []PullRequestInfo `json:"values"`
}

// PullRequestInfo struct
type PullRequestInfo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	Author      User      `json:"author"`
	Destination Commit    `json:"destination"`
	Source      Commit    `json:"source"`
	Description string    `json:"description"`
}

// PullRequestDetail struct
type PullRequestDetail struct {
	PullRequestInfo
	Participants []Reviewer `json:"participants"`
}

// Reviewer struct
type Reviewer struct {
	User     User `json:"user"`
	Approved bool `json:"approved"`
}

// User struct
type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// Commit struct
type Commit struct {
	Branch Branch `json:"branch"`
}

// Branch struct
type Branch struct {
	Name string `json:"name"`
}

// Repository struct
type Repository struct {
	Org  string
	Name string
}

// UserCredential struct
type UserCredential struct {
	Username string
	Password string
}

// IsApproved checks the pull request has been approved by
// user with the specified username
func (pr PullRequestDetail) IsApproved(username string) bool {
	for _, reviewer := range pr.Participants {
		if reviewer.Approved && reviewer.User.Username == username {
			return true
		}
	}
	return false
}

// ToString returns the description of a pull request
func (pr PullRequestInfo) ToString() string {
	loc, _ := time.LoadLocation("Local")
	updatedUtc := pr.UpdatedOn.In(loc)
	return fmt.Sprintf("%d %s %s\t%s->%s %s\n",
		pr.ID,
		updatedUtc.Format("2006-01-02 15:04"),
		pr.Author.DisplayName,
		pr.Source.Branch.Name,
		pr.Destination.Branch.Name,
		pr.Title)
}

// ToString returns the description of a pull request
func (pr PullRequestDetail) ToString() string {
	loc, _ := time.LoadLocation("Local")
	updatedUtc := pr.UpdatedOn.In(loc)
	return fmt.Sprintf("%d %s %s\t%s->%s %s\n",
		pr.ID,
		updatedUtc.Format("2006-01-02 15:04"),
		pr.Author.DisplayName,
		pr.Source.Branch.Name,
		pr.Destination.Branch.Name,
		pr.Title)
}

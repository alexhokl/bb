package models

import (
	"time"
)

// User struct
type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// CommitBranch struct
type CommitBranch struct {
	Branch Branch `json:"branch"`
}

// Commit struct
type Commit struct {
	CommitBranch
	Commit CommitInfo `json:"commit"`
}

// CommitInfo struct
type CommitInfo struct {
	Hash    string     `json:"hash"`
	Summary RawContent `json:"summary"`
}

// Branch struct
type Branch struct {
	Name string `json:"name"`
}

// PullRequestActivity struct
type PullRequestActivity struct {
	Update  Update  `json:"update,omitempty"`
	Comment Comment `json:"comment,omitempty"`
}

// Update struct
type Update struct {
	Date   time.Time `json:"date"`
	Source Commit    `json:"source"`
	Author User      `json:"author"`
}

// Comment struct
type Comment struct {
	Content   RawContent `json:"content"`
	CreatedOn time.Time  `json:"created_on"`
	UpdatedOn time.Time  `json:"updated_on"`
	User      User       `json:"user"`
}

// RawContent struct
type RawContent struct {
	Raw string `json:"raw"`
}

// Repository struct
type Repository struct {
	Org  string
	Name string
}

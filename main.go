package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type PullRequestList struct {
	PageLen int               `json:"pagelen"`
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Items   []PullRequestInfo `json:"values"`
}

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

type PullRequest struct {
	PullRequestInfo
	Participants []Reviewer `json:"participants"`
}

type Reviewer struct {
	User     User `json:"user"`
	Approved bool `json:"approved"`
}

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type Commit struct {
	Branch Branch `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}

type Remote struct {
	Org  string
	Repo string
}

type Command struct {
	Command           Commands
	PullRequestNumber int
}

type Commands int

const (
	listCommand Commands = iota
	describeCommand
)

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}

	command, errCommand := parseCommands(os.Args)
	if errCommand != nil {
		help()
		return
	}

	username, password, errCred := getCredentials()
	dumpError(errCred)

	remote, errRemote := getRemote()
	dumpError(errRemote)

	if command.Command == listCommand {
		list(remote, username, password)
	} else if command.Command == describeCommand {
		describe(remote, username, password, command.PullRequestNumber)
	}
}

func parseCommands(args []string) (*Command, error) {
	if len(args) == 1 {
		help()
	}
	if len(args) > 1 {
		if args[1] == "list" {
			return &Command{listCommand, -1}, nil
		}
		if args[1] == "describe" {
			prNumStr := args[2]
			prNum, err := strconv.Atoi(prNumStr)
			dumpError(err)
			return &Command{describeCommand, prNum}, nil
		}
	}
	return nil, errors.New("Unknown command")
}

func help() {
	fmt.Println("Here are the commands available")
	fmt.Println("- list")
	fmt.Println("- describe")
}

func list(remote *Remote, username string, password string) {
	prList, err := getPullRequestList(remote.Org, remote.Repo, username, password)
	dumpError(err)

	for _, pr := range prList.Items {
		prInfo, _ := getPullRequest(remote.Org, remote.Repo, username, password, pr.ID)
		isApproved := isApproved(prInfo, username)
		if isApproved {
			color.Cyan(toString(&pr))
		} else {
			color.Red(toString(&pr))
		}
	}
}

func describe(remote *Remote, username string, password string, pullRequestNumber int) {
	pr, err := getPullRequest(remote.Org, remote.Repo, username, password, pullRequestNumber)
	dumpError(err)

	isApproved := isApproved(pr, username)
	if isApproved {
		color.Cyan(toInfoString(pr))
	} else {
		color.Red(toInfoString(pr))
	}

	for _, reviewer := range pr.Participants {
		if reviewer.Approved {
			fmt.Printf("Approved by %s\n", reviewer.User.DisplayName)
		}
	}
}

func getCredentials() (string, string, error) {
	username := os.Getenv("bbuser")
	password := os.Getenv("bbpassword")

	if username == "" {
		return "", "", errors.New("bbuser is not set")
	}
	if password == "" {
		return "", "", errors.New("bbpassword is not set")
	}

	return username, password, nil
}

func getPullRequestList(org string, repo string, username string, password string) (*PullRequestList, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests", org, repo), nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	dumpError(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		dumpResponse(resp)
		return nil, errors.New("Failed response")
	}
	return parsePullRequestListResponse(resp), nil
}

func getPullRequest(org string, repo string, username string, password string, pullRequestNumber int) (*PullRequest, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests/%d", org, repo, pullRequestNumber), nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	dumpError(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		dumpResponse(resp)
		return nil, errors.New("Failed response")
	}
	return parsePullRequestResponse(resp), nil
}

func parsePullRequestListResponse(resp *http.Response) *PullRequestList {
	var prList PullRequestList
	errDecode := json.NewDecoder(resp.Body).Decode(&prList)
	dumpError(errDecode)
	return &prList
}

func parsePullRequestResponse(resp *http.Response) *PullRequest {
	var pr PullRequest
	errDecode := json.NewDecoder(resp.Body).Decode(&pr)
	dumpError(errDecode)
	return &pr
}

func dumpResponse(resp *http.Response) {
	_, err := io.Copy(os.Stdout, resp.Body)
	dumpError(err)
}

func dumpError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func isApproved(pr *PullRequest, username string) bool {
	for _, reviewer := range pr.Participants {
		if reviewer.Approved && reviewer.User.Username == username {
			return true
		}
	}
	return false
}

func toString(pr *PullRequestInfo) string {
	return fmt.Sprintf("%d %s %s\t%s->%s %s\n",
		pr.ID,
		pr.UpdatedOn.Format("2006-01-02 15:04"),
		pr.Author.DisplayName,
		pr.Source.Branch.Name,
		pr.Destination.Branch.Name,
		pr.Title)
}

func toInfoString(pr *PullRequest) string {
	return fmt.Sprintf("%d %s %s\t%s->%s %s\n",
		pr.ID,
		pr.UpdatedOn.Format("2006-01-02 15:04"),
		pr.Author.DisplayName,
		pr.Source.Branch.Name,
		pr.Destination.Branch.Name,
		pr.Title)
}

func getRemote() (*Remote, error) {
	cmdName := "git"
	cmdArgs := []string{"remote", "get-url", "origin"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	dumpError(err)

	remote := string(cmdOut)

	if !strings.Contains(remote, "bitbucket.org") {
		return nil, errors.New("Only BitBucket remote is supported")
	}

	remote = strings.Replace(remote, "https://bitbucket.org/", "", -1)
	remote = strings.Replace(remote, ".git", "", -1)
	remote = strings.Replace(remote, "\n", "", -1)

	parts := strings.Split(remote, "/")

	r := Remote{
		Org:  parts[0],
		Repo: parts[1],
	}

	return &r, nil
}

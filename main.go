package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
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

func main() {
	username := os.Getenv("bbuser")
	password := os.Getenv("bbpassword")

	if username == "" {
		log.Fatalln("bbuser is not set")
		return
	}
	if password == "" {
		log.Fatalln("bbpassword is not set")
		return
	}

	remote, errRemote := getRemote()
	dumpError(errRemote)

	fmt.Println(remote.Org, remote.Repo)
	prList, err := getPullRequestList(remote.Org, remote.Repo, username, password)
	dumpError(err)

	for _, pr := range prList.Items {
		fmt.Printf("%d %s %s->%s %s\n", pr.ID, pr.Author.DisplayName, pr.Source.Branch.Name, pr.Destination.Branch.Name, pr.Title)
	}
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

func parsePullRequestListResponse(resp *http.Response) *PullRequestList {
	var prList PullRequestList
	errDecode := json.NewDecoder(resp.Body).Decode(&prList)
	dumpError(errDecode)
	return &prList
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

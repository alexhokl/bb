package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// Checkout executes git checkout commnad
func Checkout(branchName string) (string, error) {
	args := []string{"checkout", branchName}
	return execute(args)
}

// GetOriginURL executes git remote command to retrieve origin URL
func GetOriginURL() (string, error) {
	args := []string{"remote", "get-url", "origin"}
	return execute(args)
}

// GetStatus executes git status command
func GetStatus() (string, error) {
	args := []string{"status", "-s"}
	return execute(args)
}

// Pull executes git pull command
func Pull() (string, error) {
	args := []string{"pull"}
	return execute(args)
}

// Fetch executes git fetch command
func Fetch() (string, error) {
	args := []string{"fetch"}
	return execute(args)
}

// GetCurrentBranchName executes git symbolic-ref to retrieve current branch name
func GetCurrentBranchName() (string, error) {
	args := []string{"symbolic-ref", "--short", "HEAD"}
	output, err := execute(args)
	if err != nil {
		return output, err
	}
	name := strings.Replace(output, "\n", "", -1)
	return name, nil
}

// DeleteBranch executes git branch command to delete a branch
func DeleteBranch(branchName string) (string, error) {
	args := []string{"branch", "-D", branchName}
	return execute(args)
}

// Merge executes git merge command to merge from a branch
func Merge(branchName string) (string, error) {
	args := []string{"merge", "--no-edit", branchName}
	return execute(args)
}

// Difftool executes git difftool command
func Difftool(branchName string) error {
	args := []string{"difftool", branchName}
	_, err := execute(args)
	return err
}

// DiffStat executes git diff to retrieve diff stat
func DiffStat(branchName string) (string, error) {
	args := []string{"diff", "--stat", branchName}
	return execute(args)
}

// GetBranchCommitComments executes git log command to retrieve branch commit comments
func GetBranchCommitComments(sourceBranchName string, destinationBranchName string) (string, error) {
	branches := fmt.Sprintf("%s..%s", destinationBranchName, sourceBranchName)
	args := []string{"log", branches, "--no-merges", "--pretty=format:'%s %b'"}
	return execute(args)
}

func execute(args []string) (string, error) {
	byteOutput, err := exec.Command("git", args...).Output()
	return string(byteOutput), err
}

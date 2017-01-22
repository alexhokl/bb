package git

import "os/exec"

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
	return execute(args)
}

// DeleteBranch execute git branch command to delete a branch
func DeleteBranch(branchName string) (string, error) {
	args := []string{"branch", "-D", branchName}
	return execute(args)
}

func execute(args []string) (string, error) {
	byteOutput, err := exec.Command("git", args...).Output()
	return string(byteOutput), err
}

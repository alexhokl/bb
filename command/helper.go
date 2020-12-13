package command

import (
	"runtime"
)

func getOpenCommand(args ...string) (string, []string) {
	switch runtime.GOOS {
	case "windows":
		cmdArgs := []string{"/C", "start"}
		for _, a := range args {
			cmdArgs = append(cmdArgs, a)
		}
		return "cmd", cmdArgs
	case "darwin":
		return "open", args
	default:
		return "xdg-open", args
	}
}

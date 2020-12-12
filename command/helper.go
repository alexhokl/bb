package command

import (
	"encoding/json"
	"fmt"
	"os"
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

func writeJSONFile(path string, obj interface{}) error {
	if path == "" {
		return fmt.Errorf("path is not specified")
	}
	if obj == nil {
		return fmt.Errorf("obj cannot be empty")
	}

	if _, err := os.Stat(path); os.IsExist(err) {
		errRemove := os.Remove(path)
		if errRemove != nil {
			return errRemove
		}
	}

	file, errOpen := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if errOpen != nil {
		return errOpen
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err := encoder.Encode(obj)
	return err
}

func getDistinctNames(names []string) []string {
	m := map[string]bool{}

	for v := range names {
		m[names[v]] = true
	}

	d := []string{}
	for k := range m {
		d = append(d, k)
	}
	return d
}

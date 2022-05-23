package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func checkRequirements() error {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatal("kubectl not found in $PATH!\nPlease install kubectl from https://kubernetes.io/de/docs/tasks/tools/install-kubectl/\n")
		return err
	}

	log.Printf("Found kubectl at %s.", path)

	return nil
}

func getContextSliceIndexFromContext(context Context, contextSlice []Context) int {
	for index, value := range contextSlice {
		if value.Name == context.Name {
			return index
		}
	}

	return -1
}

func getExecutable() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return executable
}

// opens [url] using the `open` command
func openURL(url string) error {
	return exec.Command("open", url).Start()
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

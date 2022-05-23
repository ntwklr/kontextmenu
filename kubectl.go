package main

import (
	"log"
    "os/exec"
	"strings"
)

func getContexts() []Context {
	output, err := exec.Command("kubectl", "config", "get-contexts", "--no-headers", "-oname").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	contexts := []Context{}

	for _, item := range lines {
		log.Printf("Read context \"%s\".", item)

		contexts = append(contexts, Context {
			Name: item,
		})
	}

	return contexts
}

func getCurrentContext() Context {
	output, err := exec.Command("kubectl", "config", "current-context").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return Context{Name: strings.TrimSpace(string(output))}
}

func (a *App) useContext(context Context) error {
	output, err := exec.Command("kubectl", "config", "use-context", context.Name).CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Print(strings.TrimSpace(string(output)))

	a.currentContext = getCurrentContext()

	return nil
}
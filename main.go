package main

import (
	"log"
	"os"
	"path"

	"github.com/emersion/go-autostart"
    "github.com/getlantern/systray"
)

const aboutURL = "https://github.com/ntwklr/kontextmenu"

type Context struct {
	Name 	string
}

type App struct {
	autostart 		autostart.App

	contexts		[]Context

	currentContext 	Context

	isSystrayReady 	bool
}

func main() {
	err := checkRequirements()
	if err != nil {
		log.Fatalf("Error while checking requirements: %s", err.Error())
		os.Exit(1)
	}

	app := App {
		autostart: 		autostart.App {
							Name: "kontextmenu",
							DisplayName: "Kubernetes context switcher",
							Exec: []string{path.Dir(getExecutable())},
						},

		contexts: 		getContexts(),

		currentContext: getCurrentContext(),

		isSystrayReady: false,
	}

	systray.Run(func() {
		app.onSystrayReady()
	}, func() {
		app.onSystrayExit()
	})
}
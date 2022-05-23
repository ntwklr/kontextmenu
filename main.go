package main

import (
	"embed"
	"log"
	"os"
	"path"

	"github.com/emersion/go-autostart"
	"github.com/getlantern/systray"
)

//go:embed assets/*
var fs embed.FS

const aboutURL = "https://github.com/ntwklr/kontextmenu"

type Context struct {
	Name string
}

type App struct {
	icon                 []byte
	iconTemplate         []byte
	iconTemplateFallback []byte

	autostart autostart.App

	contexts []Context

	currentContext Context

	isSystrayReady bool
}

func main() {
	err := checkRequirements()
	if err != nil {
		log.Fatalf("Error while checking requirements: %s", err.Error())
		os.Exit(1)
	}

	icon, _ := fs.ReadFile("assets/kubernetes-1024.png")
	iconTemplate, _ := fs.ReadFile("assets/kubernetes-1024_template.png")
	iconTemplateFallback, _ := fs.ReadFile("assets/kubernetes-1024.ico")

	app := App{
		icon:                 icon,
		iconTemplate:         iconTemplate,
		iconTemplateFallback: iconTemplateFallback,

		autostart: autostart.App{
			Name:        "kontextmenu",
			DisplayName: "Kubernetes context switcher",
			Exec:        []string{path.Dir(getExecutable())},
		},

		contexts: getContexts(),

		currentContext: getCurrentContext(),

		isSystrayReady: false,
	}

	systray.Run(func() {
		app.onSystrayReady()
	}, func() {
		app.onSystrayExit()
	})
}

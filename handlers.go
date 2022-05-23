package main

import (
	"log"
    "github.com/getlantern/systray"
)

func (a *App) handleContextSelected(mContexts []*systray.MenuItem, index int) {
	prevIndex := getContextSliceIndexFromContext(a.currentContext, a.contexts[:])

	if prevIndex >= 0 && prevIndex < len(mContexts) {
		mContexts[prevIndex].Uncheck()
	}
	mContexts[index].Check()

	a.useContext(a.contexts[index])
}

func (a *App) handleLaunchAtLoginClicked(item *systray.MenuItem) {
	if a.autostart.IsEnabled() {
		item.Uncheck()
		log.Print("Remove kontexmenu from Autostart.")

		if err := a.autostart.Disable(); err != nil {
			log.Fatal(err)
		}
	} else {
		item.Check()
		log.Print("Add kontexmenu to Autostart.")

		if err := a.autostart.Enable(); err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) handleAboutClicked() {
	openURL(aboutURL)
}

func (a *App) handleQuitClicked() {
	systray.Quit()
}
package main

import (
	"github.com/getlantern/systray"
)

func (a *App) onSystrayReady() {
	systray.SetIcon(getIcon("assets/kubernetes-1024.png"))
	systray.SetTemplateIcon(getIcon("assets/kubernetes-1024_template.png"), getIcon("assets/kubernetes-1024.ico"))
	//systray.SetTitle("Awesome App")
	systray.SetTooltip("Kubernetes context switcher")

	createContextSelectionItems := func(
		selectedItem Context,
	) (
		itemSelected chan int, menuItems []*systray.MenuItem,
	) {
		selectedItemIndex := getContextSliceIndexFromContext(selectedItem, a.contexts[:])

		for index, item := range a.contexts {
			menuItems = append(
				menuItems,
				systray.AddMenuItemCheckbox(
					item.Name,
					"",
					index == selectedItemIndex,
				),
			)
		}

		itemSelected = make(chan int)
		for i, mItem := range menuItems {
			go func(c chan struct{}, index int) {
				for range c {
					itemSelected <- index
				}
			}(mItem.ClickedCh, i)
		}

		return itemSelected, menuItems
	}

	contextSelected, mContexts := createContextSelectionItems(a.currentContext)

	systray.AddSeparator()

	mLaunchAtLogin := systray.AddMenuItemCheckbox(
		"Launch at Login",
		"",
		a.autostart.IsEnabled(),
	)

	systray.AddSeparator()

	mAbout := systray.AddMenuItem("About", "")
	mQuit := systray.AddMenuItem("Quit kontextmenu", "")

	go func() {
		for {
			select {
			case index := <-contextSelected:
				a.handleContextSelected(mContexts, index)
			case <-mLaunchAtLogin.ClickedCh:
				a.handleLaunchAtLoginClicked(mLaunchAtLogin)
			case <-mAbout.ClickedCh:
				a.handleAboutClicked()
			case <-mQuit.ClickedCh:
				a.handleQuitClicked()
			}
		}
	}()

	a.isSystrayReady = true
}

func (a *App) onSystrayExit() {
	// clean up here
}

package main

import (
	"fmt"
	"os"

	"github.com/pkg/browser"
	hook "github.com/robotn/gohook"

	"github.com/itchyny/volume-go"

	starter "github.com/emersion/go-autostart"

	"github.com/platilich/randomick"
)

func main() {
	fmt.Println("started (command + x)\nexit (cotrol + c)")
	autostart()

	hook.Register(hook.KeyDown, []string{"command", "x"}, HideMyScreen)

	s := hook.Start()
	<-hook.Process(s)

}

func HideMyScreen(_ hook.Event) {
	web := randomick.Choice("https://youtube.com", "https://google.com", "https://wikipedia.org")

	browser.OpenURL(web)
	volume.Mute()

}

func autostart() {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	// init autostart
	app := &starter.App{
		Name:        "focusback", // name plist file
		DisplayName: "FocusBack Application",
		Exec:        []string{execPath}, // path to programm
	}

	if app.IsEnabled() {
		return
	}

	err = app.Enable()
	if err != nil {
		fmt.Println(err)

	}

}

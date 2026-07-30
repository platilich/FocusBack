package main

import (
	"fmt"

	"github.com/pkg/browser"
	hook "github.com/robotn/gohook"

	"github.com/itchyny/volume-go"
)

func main() {
	fmt.Println("started (command + 0)\nexit (cotrol + c)")

	hook.Register(hook.KeyDown, []string{"command", "0"}, HideMyScreen)

	s := hook.Start()
	<-hook.Process(s)

}

func HideMyScreen(_ hook.Event) {
	browser.OpenURL("https://ru.wikipedia.org/wiki/")
	volume.Mute()

}

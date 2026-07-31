package main

import (
	"fmt"

	"github.com/pkg/browser"
	hook "github.com/robotn/gohook"

	"github.com/itchyny/volume-go"

	"os"

	starter "github.com/emersion/go-autostart"
)

func main() {
	fmt.Println("started (command + 0)\nexit (cotrol + c)")
	autostart()

	hook.Register(hook.KeyDown, []string{"command", "0"}, HideMyScreen)

	s := hook.Start()
	<-hook.Process(s)

}

func HideMyScreen(_ hook.Event) {
	browser.OpenURL("https://ru.wikipedia.org/wiki/")
	volume.Mute()

}

func autostart() {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	// 2. Инициализируем автозапуск
	app := &starter.App{
		Name:        "focusback",             // Имя для plist-файла
		DisplayName: "FocusBack Application", // Читаемое имя приложения
		Exec:        []string{execPath},      // Путь к программе
	}

	if app.IsEnabled() {
		return
	}

	err = app.Enable()
	if err != nil {
		fmt.Println(err)

	}

}

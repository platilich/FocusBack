# FocusBack

## Written in Go. Simple and minimalist. Distributed under the MIT License.

### FocusBack is a small utility that helps you quickly switch your attention back when someone enters your room or you need to hide distractions.

## How to use it?

It's simple: just press:

```
Command + 0
```

The app will open your chosen website and mute the volume.



## Installation
### WARNING There will be a proper installation process here soon; I'll add releases shortly, so you won't have to do all this.


Step 1:

```
git clone https://github.com/platilich/FocusBack.git
```

Step 2:

```
go build main.go
```


Step 3:

```
sudo ./main
```




## Want to change the website?

No problem.

Open `main.go` and change the URL on line 28:

```go
browser.OpenURL("https://your_web_site")
```

## Want to change the hotkey?

Modify this line on line 20:

```go
hook.Register(hook.KeyDown, []string{"your_first_key", "second_key"}, HideMyScreen)
```

Replace the keys with your preferred shortcut.

## License

MIT License.

# Sysdeep atk

Another Golang Tcl/Tk binding GUI ToolKit

    go get github.com/sysdeep/atk

Go docs: https://pkg.go.dev/github.com/visualfc/atk/tk

Original repo - https://github.com/visualfc/atk
Another tk9 - https://gitlab.com/cznic/tk9.0

### Install Tcl/Tk

http://www.tcl-lang.org

- MacOS X

  https://www.activestate.com/activetcl/downloads

- Windows

  https://www.activestate.com/activetcl/downloads

  https://github.com/visualfc/tcltk_mingw

- Ubuntu

  $ sudo apt install tk-dev

- CentOS

  $ sudo yum install tk-devel

### Demo

see examples

### Sample

```go
package main

import (
	"github.com/sysdeep/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := &Window{tk.RootWindow()}
	lbl := tk.NewLabel(mw, "Hello ATK")
	btn := tk.NewButton(mw, "Quit")
	btn.OnCommand(func() {
		tk.Quit()
	})
	tk.NewVPackLayout(mw).AddWidgets(lbl, tk.NewLayoutSpacer(mw, 0, true), btn)
	mw.ResizeN(300, 200)
	return mw
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("ATK Sample")
		mw.Center(nil)
		mw.ShowNormal()
	})
}
```

## TODO

- [ ] label
- [ ] toplevel

  - [ ] wm iconname $w "label"

- [ ] options
  - [ ] сейчас простая экранизация {} - но она не всегда подходит
- [ ] common
  - [ ] единицы размеров не только в цифрах но ещё и в i, p, m...
  - [ ] pack - принимает несколько элементов
  - [ ] вызов виджетов сделать более каноничным(см. tk9)

```go

// command id options
// label $w.left.l1 -text "First label"

// example
tk.Label(parent, tk.LabelOpts().Text(""))
tk.TLabel(parent, tk.TLabelOpts().Text(""))

// LabelOpts - ограничитель-построитель опций только для Label
```

## Devel

### Fonts

```wish
font configure TkDefaultFont -size 18
font names
font families
font configure TkDefaultFont -family {Noto Looped Thai}
```

package main

import (
	"github.com/sysdeep/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {

	// tk.SetMainTheme(nil)
	// fmt.Println("use theme?")
	// fmt.Println(tk.HasTheme())
	// fmt.Println(tk.MainTheme())

	w := tk.RootWindow()

	lbl_1 := tk.NewLabel(w, "first")
	tk.Pack(lbl_1, tk.PackAttrSideTop())

	// close
	btnClose := tk.NewButton(w, "Close")
	btnClose.OnCommand(func() {
		w.Destroy()
	})
	tk.Pack(btnClose, tk.PackAttrSide(tk.SideBottom))

	// bind
	w.BindEvent("<Control-q>", func(e *tk.Event) {
		w.Destroy()
	})
	return &Window{w}
}

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.SetTitle("Label Demo")
		w.Center(nil)
		w.ResizeN(400, 300)
		w.ShowNormal()
	})
}

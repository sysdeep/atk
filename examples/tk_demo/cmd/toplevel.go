package main

import "github.com/sysdeep/atk/tk"

func MakeToplevel(cb func(tk.Widget) (tk.Widget, string)) {
	w := tk.NewWindow()

	content, title := cb(w)

	tk.Pack(content, tk.PackAttrSideTop(), tk.PackAttrExpand(true), tk.PackAttrFillBoth())

	// controls
	makeControls(w)

	w.BindEvent("<Control-q>", func(e *tk.Event) {
		w.Destroy()
	})

	w.SetTitle(title)
	w.ResizeN(800, 600)

	// TODO
	// wm iconname $w "label"
	w.ShowNormal()
}

func makeControls(parent tk.Widget) {
	// bottom
	controlsFrame := tk.NewFrame(parent)
	tk.Pack(controlsFrame, tk.PackAttrSideBottom(), tk.PackAttrFillX())

	exitBtn := tk.NewButton(controlsFrame, "Exit")
	exitBtn.OnCommand(func() {
		parent.Destroy()
	})
	tk.Pack(exitBtn, tk.PackAttrSideRight())

}

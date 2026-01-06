package common

import "github.com/sysdeep/atk/tk"

func MakeApp(cb func(tk.Widget) (tk.Widget, string)) {
	w := tk.RootWindow()

	content, title := cb(w)

	tk.Pack(content, tk.PackAttrSideTop(), tk.PackAttrExpand(true), tk.PackAttrFillBoth())

	// controls
	MakeControls(w)

	w.BindEvent("<Control-q>", func(e *tk.Event) {
		w.Destroy()
	})

	w.SetTitle(title)
	w.ResizeN(800, 600)

	// TODO
	// wm iconname $w "label"
	w.ShowNormal()
}

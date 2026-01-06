package main

import "github.com/sysdeep/atk/tk"

func doLabels1() {
	w := tk.NewWindow()

	intro := `Five labels are displayed below: three textual ones on the left, and an image label and a text label on the right.  Labels are pretty boring because you can't do anything with them.`
	tk.NewLabel(w, intro, tk.LabelOptWrapLength(344), tk.LabelOptJustify("left")).Pack(tk.PackAttrSideTop())

	// bottom
	controlsFrame := tk.NewFrame(w)
	tk.Pack(controlsFrame, tk.PackAttrSideBottom(), tk.PackAttrFillX())

	exitBtn := tk.NewButton(controlsFrame, "Exit")
	exitBtn.OnCommand(func() {
		w.Destroy()
	})
	tk.Pack(exitBtn, tk.PackAttrSideRight())

	// bind
	w.BindEvent("<Control-q>", func(e *tk.Event) {
		w.Destroy()
	})

	w.SetTitle("Label Demonstration")
	w.ResizeN(800, 600)

	// TODO
	// wm iconname $w "label"
	w.ShowNormal()
}

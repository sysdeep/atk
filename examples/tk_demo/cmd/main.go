package main

import (
	"github.com/sysdeep/atk/examples/tk_demo/label"
	"github.com/sysdeep/atk/examples/tk_demo/unicode"
	"github.com/sysdeep/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		w := tk.RootWindow()

		f := tk.NewFrame(w)
		tk.Pack(f, tk.PackAttrSideTop(), tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

		frames := [][]*tk.LabelFrame{
			{
				makeFrameLabels(f),
			},
		}

		for i, row := range frames {
			for j, fr := range row {
				tk.Grid(fr, tk.GridAttrColumn(j), tk.GridAttrRow(i))
			}
		}

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

		w.SetTitle("Std TK Deno")
		w.Center(nil)
		w.ResizeN(800, 600)
		w.ShowNormal()
	})
}

func makeFrameLabels(parent tk.Widget) *tk.LabelFrame {
	frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Labels"))

	b1 := tk.NewButton(frame, label.LabelTitle)
	b1.OnCommand(doLabel)
	b1.Pack()

	b2 := tk.NewButton(frame, unicode.UnicodeTitle)
	b2.OnCommand(doUnicode)
	b2.Pack()

	return frame
}

func doLabel() {
	MakeToplevel(func(parent tk.Widget) (tk.Widget, string) {
		return label.NewLabel(parent), label.LabelTitle
	})
}

func doUnicode() {
	MakeToplevel(func(parent tk.Widget) (tk.Widget, string) {
		return unicode.NewUnicode(parent), unicode.UnicodeTitle
	})
}

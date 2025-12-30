package main

import (
	"github.com/sysdeep/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		w := tk.RootWindow()
		w.SetTitle("Label Demo")
		w.Center(nil)
		w.ResizeN(800, 600)

		tk.NewLabelOptions().Text("test").
			Background("green").
			PadY(32).
			PadX(44).Foreground("red").
			Make(w).Pack()

		lbl2 := tk.NewLabelOptions().Text("test configure").Make(w)
		lbl2.Pack()
		lbl2.Configure(tk.NewLabelOptions().Background("blue").GetOptions()...)

		lbl3 := tk.NewLabelOptions().Text("test configure2").Make(w)
		lbl3.Pack()
		tk.NewLabelOptions().Background("black").Foreground("white").Configure(lbl3)

		w.ShowNormal()
	})
}

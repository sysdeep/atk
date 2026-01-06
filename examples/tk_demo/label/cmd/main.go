package main

import (
	"github.com/sysdeep/atk/examples/tk_demo/common"
	"github.com/sysdeep/atk/examples/tk_demo/label"
	"github.com/sysdeep/atk/tk"
)

func main() {
	tk.MainLoop(func() {

		common.MakeApp(func(parent tk.Widget) (tk.Widget, string) {
			return label.NewLabel(parent), "Label Demonstration"
		})

	})
}

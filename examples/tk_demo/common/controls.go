package common

import "github.com/sysdeep/atk/tk"

func MakeControls(parent tk.Widget) {
	// bottom
	controlsFrame := tk.NewFrame(parent)
	tk.Pack(controlsFrame, tk.PackAttrSideBottom(), tk.PackAttrFillX())

	exitBtn := tk.NewButton(controlsFrame, "Exit")
	exitBtn.OnCommand(func() {
		parent.Destroy()
	})
	tk.Pack(exitBtn, tk.PackAttrSideRight())

}

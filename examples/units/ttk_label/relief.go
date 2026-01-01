package main

import "github.com/sysdeep/atk/tk"

func makeReliefFrame(parent tk.Widget) *tk.LabelFrame {
	frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Relief"))

	var borderStyleName = []string{"flat", "groove", "raised", "ridge", "solid", "sunken"}

	for i, s := range borderStyleName {
		lbl := tk.NewLabel(frame, "Relief: "+s)
		lbl.Configure(
			tk.LabelOptRelief(s),
			tk.LabelOptWidth(24),
			tk.LabelOptPadY(4),
		)

		tk.Grid(lbl, tk.GridAttrRow(i), tk.GridAttrColumn(0))
	}

	return frame
}

package main

import (
	"github.com/sysdeep/atk/tk"
	"github.com/sysdeep/atk/tk/enum/cursor"
)

func makeCursorFrame(parent tk.Widget) *tk.LabelFrame {
	frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Cursor"))

	var sursors = []string{cursor.Boat, cursor.UrAngle, cursor.BottomTee, cursor.Circle}

	for i, s := range sursors {
		lbl := tk.NewLabel(frame, "Cursor: "+s)
		lbl.Configure(
			tk.LabelOptCursor(s),
			tk.LabelOptWidth(24),
			tk.LabelOptPadY(4),
		)

		tk.Grid(lbl, tk.GridAttrRow(i), tk.GridAttrColumn(0))
	}

	return frame
}

package main

import (
	"fmt"

	"github.com/sysdeep/atk/tk"
	"github.com/sysdeep/atk/tk/enum/color"
)

func makeColorFrame(parent tk.Widget) *tk.LabelFrame {
	frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Color"))

	colors := [][]string{
		{color.Antiquewhite, color.Gray0},
		{color.Green, color.Yellow},
	}

	for _, row := range colors {

		tk.NewLabel(frame, fmt.Sprintf("BG %s and FG %s", row[0], row[1]),
			tk.LabelOptBackground(row[0]),
			tk.LabelOptForeground(row[1]),
		).Pack()
	}

	var LabelStates = []string{tk.LabelStateNormal, tk.LabelStateActive, tk.LabelStateDisabled}
	for _, state := range LabelStates {
		tk.NewLabel(frame, fmt.Sprintf("State: %s", state),
			tk.LabelOptState(state),
		).Pack()
	}

	return frame
}

package main

import "github.com/sysdeep/atk/tk"

func makeFontsFrame(parent tk.Widget) *tk.LabelFrame {
	frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Fonts"))

	var sysFontNameList = []string{
		// "TKDefaultFont",
		// "TKTextFont",
		// "TKFixedFont",
		// "TKMenuFont",
		// "TKHeadingFont",
		// "TKCaptionFont",
		// "TKSmallCaptionFont",
		// "TKIconFont",
		// "TKTooltipFont",

		"Gubbi",
	}

	for i, txt := range sysFontNameList {
		lbl := tk.NewLabel(frame, txt)
		lbl.Configure(tk.LabelOptionFont(txt))
		tk.Grid(lbl, tk.GridAttrColumn(0), tk.GridAttrRow(i))
	}

	return frame
}

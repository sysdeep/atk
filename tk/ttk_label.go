// https://www.tcl-lang.org/man/tcl8.6/TkCmd/ttk_label.htm

package tk

/*
NAME
ttk::label — Display a text string and/or image
STANDARD OPTIONS
-padding, padding, Padding
-style, style, Style
-textvariable, textVariable, Variable
WIDGET-SPECIFIC OPTIONS
-anchor, anchor, Anchor
-background, frameColor, FrameColor
-font, font, Font
-foreground, textColor, TextColor
-justify, justify, Justify
-relief, relief, Relief
-wraplength, wrapLength, WrapLength
WIDGET COMMAND
*/
// ttk::label
type TLabel struct {
	BaseWidget
}

func NewTLabel(parent Widget, text string, options ...OptionAdapter) *TLabel {

	iid := makeNamedWidgetId(parent, "ttk_label")

	options = append(options, TLabelOptText(text))

	info := CreateWidgetInfoOptions(iid, "ttk::label", options)

	w := &TLabel{
		NewBaseWidget(iid, &info),
	}

	RegisterWidget(w)
	return w
}

// new options api ------------------------------------------------------------
// TODO: move to base widget

// option aliases -------------------------------------------------------------
var TLabelOptText = OptText
var TLabelOptUnderline = OptUnderline
var TLabelOptWidth = OptWidth
var TLabelOptCompound = OptCompound
var TLabelOptImage = OptImage
var TLabelOptTakeFocus = OptTakeFocus
var TLabelOptCursor = OptCursor
var TLabelOptClass = OptClass
var TLabelOptState = OptState

// label specific options -----------------------------------------------------
// TODO

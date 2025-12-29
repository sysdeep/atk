// https://www.tcl-lang.org/man/tcl8.6/TkCmd/ttk_label.htm

package tk

// ttk::label
type TLabel struct {
	BaseWidget
}

// TODO: use new options api
func NewTLabel(parent Widget, text string, options ...OptionAdapter) *TLabel {

	iid := makeNamedWidgetId(parent, "ttk_label")

	options = append(options, TLabelOptionText(text))

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
var TLabelOptionText = OptionText
var TLabelOptionUnderline = OptionUnderline

// label specific options -----------------------------------------------------
// TODO

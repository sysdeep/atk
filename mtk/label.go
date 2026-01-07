package mtk

import "github.com/sysdeep/atk/tk"

// Label — Create and manipulate 'label' non-interactive text or image widgets
//
// # Description
//
// The label command creates a new window (given by the
// pathName argument) and makes it into a label widget.
// Additional
// options, described above, may be specified on the command line
// or in the option database
// to configure aspects of the label such as its colors, font,
// text, and initial relief.  The label command returns its
// pathName argument.  At the time this command is invoked,
// there must not exist a window named pathName, but
// pathName's parent must exist.
//
// A label is a widget that displays a textual string, bitmap or image.
// If text is displayed, it must all be in a single font, but it
// can occupy multiple lines on the screen (if it contains newlines
// or if wrapping occurs because of the -wraplength option) and
// one of the characters may optionally be underlined using the
// The label can be manipulated in a few simple ways, such as
// changing its relief or text, using the commands described below.
//
// Use [Window.Label] to create a Label with a particular parent.
//
// More information might be available at the [Tcl/Tk label] page.
//
// # Standard options
//
//   - [Activebackground]
//   - [Activeforeground]
//   - [Anchor]
//   - [Background]
//   - [Bitmap]
//   - [Borderwidth]
//   - [Compound]
//   - [Cursor]
//   - [Disabledforeground]
//   - [Foreground]
//   - [Highlightbackground]
//   - [Highlightcolor]
//   - [Highlightthickness]
//   - [Image]
//   - [Justify]
//   - [Padx]
//   - [Pady]
//   - [Relief]
//   - [Takefocus]
//   - [Textvariable]
//   - [Txt]
//   - [Underline]
//   - [Wraplength]
//
// # Widget specific options
//
// [Height]
//
// Specifies a desired height for the label.
// If an image or bitmap is being displayed in the label then the value is in
// screen units (i.e. any of the forms acceptable to Tk_GetPixels);
// for text it is in lines of text.
// If this option is not specified, the label's desired height is computed
// from the size of the image or bitmap or text being displayed in it.
//
// [State]
//
// Specifies one of three states for the label:  normal, active,
// or disabled.  In normal state the button is displayed using the
// -foreground and -background options.  In active state
// the label is displayed using the -activeforeground and
// -activebackground options.  In the disabled state the
// -disabledforeground and -background options determine how
// the button is displayed.
//
// [Width]
//
// Specifies a desired width for the label.
// If an image or bitmap is being displayed in the label then the value is in
// screen units (i.e. any of the forms acceptable to Tk_GetPixels);
// for text it is in characters.
// If this option is not specified, the label's desired width is computed
// from the size of the image or bitmap or text being displayed in it.
//
// [Tcl/Tk label]: https://www.tcl-lang.org/man/tcl9.0/TkCmd/label.html
func Label(parent tk.Widget, options ...Opt) *tk.Label {
	return tk.NewLabel(parent, "TODO")
}

// Label — Create and manipulate 'label' non-interactive text or image widgets
//
// The resulting [Window] is a child of 'w'
//
// For details please see [Label]
// func (w *Window) Label(options ...Opt) *LabelWidget {
// 	return &LabelWidget{w.newChild("label", options...)}
// }

// LabelWidget represents the Tcl/Tk label widget/window
// type LabelWidget struct {
// 	*Window
// }

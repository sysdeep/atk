// https://www.tcl-lang.org/man/tcl8.6/TkCmd/label.htm

package tk

// tk::label
type Label struct {
	BaseWidget
}

// NewLabel
func NewLabel(parent Widget, text string, options ...OptionAdapter) *Label {

	iid := makeNamedWidgetId(parent, "tk_label")

	options = append(options, LabelOptionText(text))

	info := CreateWidgetInfoOptions(iid, "label", options)

	w := &Label{
		NewBaseWidget(iid, &info),
	}

	RegisterWidget(w)
	return w
}

// option aliases -------------------------------------------------------------
var LabelOptionBackground = OptionBackground
var LabelOptionActiveBackground = OptionActiveBackground
var LabelOptionForeground = OptionForeground
var LabelOptionActiveForeground = OptionActiveForeground
var LabelOptionRelief = OptionRelief
var LabelOptionBorderWidth = OptionBorderWidth
var LabelOptionPadX = OptionPadX
var LabelOptionPadY = OptionPadY
var LabelOptionText = OptionText
var LabelOptionJustify = OptionJustify
var LabelOptionAnchor = OptionAnchor
var LabelOptionWrapLength = OptionWrapLength
var LabelOptionUnderline = OptionUnderline
var LabelOptionFont = OptionFont
var LabelOptionCompound = OptionCompound
var LabelOptionImage = OptionImage
var LabelOptionTakeFocus = OptionTakeFocus

// label specific options -----------------------------------------------------
/*
Command-Line Name: -width
Database Name: width
Database Class: Width
Specifies a desired width for the label. If an image or bitmap is being displayed in the label then the value is in screen units (i.e. any of the forms acceptable to Tk_GetPixels); for text it is in characters. If this option is not specified, the label's desired width is computed from the size of the image or bitmap or text being displayed in it.
*/
func LabelOptionWidth(value int) *intOption {
	return newIntOption("width", value)
}

/*
Command-Line Name: -height
Database Name: height
Database Class: Height
Specifies a desired height for the label. If an image or bitmap is being displayed in the label then the value is in screen units (i.e. any of the forms acceptable to Tk_GetPixels); for text it is in lines of text. If this option is not specified, the label's desired height is computed from the size of the image or bitmap or text being displayed in it.
*/
func LabelOptionHeight(value int) *intOption {
	return newIntOption("height", value)
}

/*
Command-Line Name: -state
Database Name: state
Database Class: State
Specifies one of three states for the label: normal, active, or disabled.
In normal state the button is displayed using the -foreground and -background options.
In active state the label is displayed using the -activeforeground and -activebackground options.
In the disabled state the -disabledforeground and -background options determine how the button is displayed.
*/
func LabelOptionState(state string) *stringOption {
	return newStringOption("state", state)
}

// fast configure aliases -----------------------------------------------------
func (l *Label) SetText(value string) {
	l.Configure(LabelOptionText(value))
}

// org get-set ----------------------------------------------------------------
// Зачем это?
// func (w *Label) Attach(id string) error {
// 	info, err := CheckWidgetInfo(id, WidgetTypeLabel)
// 	if err != nil {
// 		return err
// 	}
// 	w.id = id
// 	w.info = info
// 	RegisterWidget(w)
// 	return nil
// }

// func (w *Label) SetBackground(color string) error {
// 	setObjText("atk_tmp_text", color)
// 	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
// }

// func (w *Label) Background() string {
// 	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
// 	return r
// }

// func (w *Label) SetBorderWidth(width int) error {
// 	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
// }

// func (w *Label) BorderWidth() int {
// 	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
// 	return r
// }

// func (w *Label) SetForeground(color string) error {
// 	setObjText("atk_tmp_text", color)
// 	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.id))
// }

// func (w *Label) Foreground() string {
// 	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.id))
// 	return r
// }

// func (w *Label) SetReliefStyle(relief ReliefStyle) error {
// 	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
// }

// func (w *Label) ReliefStyle() ReliefStyle {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
// 	return parserReliefStyleResult(r, err)
// }

// func (w *Label) SetFont(font Font) error {
// 	if font == nil {
// 		return ErrInvalid
// 	}
// 	return eval(fmt.Sprintf("%v configure -font {%v}", w.id, font.Id()))
// }

// func (w *Label) Font() Font {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.id))
// 	return parserFontResult(r, err)
// }

// func (w *Label) SetAnchor(anchor Anchor) error {
// 	return eval(fmt.Sprintf("%v configure -anchor {%v}", w.id, anchor))
// }

// func (w *Label) Anchor() Anchor {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -anchor", w.id))
// 	return parserAnchorResult(r, err)
// }

// func (w *Label) SetJustify(justify Justify) error {
// 	return eval(fmt.Sprintf("%v configure -justify {%v}", w.id, justify))
// }

// func (w *Label) Justify() Justify {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -justify", w.id))
// 	return parserJustifyResult(r, err)
// }

// func (w *Label) SetWrapLength(wraplength int) error {
// 	return eval(fmt.Sprintf("%v configure -wraplength {%v}", w.id, wraplength))
// }

// func (w *Label) WrapLength() int {
// 	r, _ := evalAsInt(fmt.Sprintf("%v cget -wraplength", w.id))
// 	return r
// }

// func (w *Label) SetImage(image *Image) error {
// 	if image == nil {
// 		return ErrInvalid
// 	}
// 	return eval(fmt.Sprintf("%v configure -image {%v}", w.id, image.Id()))
// }

// func (w *Label) Image() *Image {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -image", w.id))
// 	return parserImageResult(r, err)
// }

// func (w *Label) SetCompound(compound Compound) error {
// 	return eval(fmt.Sprintf("%v configure -compound {%v}", w.id, compound))
// }

// func (w *Label) Compound() Compound {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -compound", w.id))
// 	return parserCompoundResult(r, err)
// }

// func (w *Label) SetText(text string) error {
// 	setObjText("atk_tmp_text", text)
// 	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.id))
// }

// func (w *Label) Text() string {
// 	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.id))
// 	return r
// }

// func (w *Label) SetWidth(width int) error {
// 	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
// }

// func (w *Label) Width() int {
// 	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
// 	return r
// }

// func (w *Label) SetPaddingN(padx int, pady int) error {
// 	if w.info.IsTtk {
// 		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
// 	}
// 	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
// }

// func (w *Label) PaddingN() (int, int) {
// 	var r string
// 	var err error
// 	if w.info.IsTtk {
// 		r, err = evalAsString(fmt.Sprintf("%v cget -padding", w.id))
// 	} else {
// 		r1, _ := evalAsString(fmt.Sprintf("%v cget -padx", w.id))
// 		r2, _ := evalAsString(fmt.Sprintf("%v cget -pady", w.id))
// 		r = r1 + " " + r2
// 	}
// 	return parserPaddingResult(r, err)
// }

// func (w *Label) SetPadding(pad Pad) error {
// 	return w.SetPaddingN(pad.X, pad.Y)
// }

// func (w *Label) Padding() Pad {
// 	x, y := w.PaddingN()
// 	return Pad{x, y}
// }

// func (w *Label) SetState(state State) error {
// 	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
// }

// func (w *Label) State() State {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
// 	return parserStateResult(r, err)
// }

// func (w *Label) SetTakeFocus(takefocus bool) error {
// 	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
// }

// func (w *Label) IsTakeFocus() bool {
// 	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
// 	return r
// }

// label attrs ----------------------------------------------------------------
// func LabelAttrBackground(color string) *WidgetAttr {
// 	return &WidgetAttr{"background", color}
// }

// func LabelAttrBorderWidth(width int) *WidgetAttr {
// 	return &WidgetAttr{"borderwidth", width}
// }

// func LabelAttrForground(color string) *WidgetAttr {
// 	return &WidgetAttr{"foreground", color}
// }

// func LabelAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
// 	return &WidgetAttr{"relief", relief}
// }

// func LabelAttrFont(font Font) *WidgetAttr {
// 	if font == nil {
// 		return nil
// 	}
// 	return &WidgetAttr{"font", font.Id()}
// }

// func LabelAttrAnchor(anchor Anchor) *WidgetAttr {
// 	return &WidgetAttr{"anchor", anchor}
// }

// func LabelAttrJustify(justify Justify) *WidgetAttr {
// 	return &WidgetAttr{"justify", justify}
// }

// func LabelAttrWrapLength(wraplength int) *WidgetAttr {
// 	return &WidgetAttr{"wraplength", wraplength}
// }

// func LabelAttrImage(image *Image) *WidgetAttr {
// 	if image == nil {
// 		return nil
// 	}
// 	return &WidgetAttr{"image", image.Id()}
// }

// func LabelAttrCompound(compound Compound) *WidgetAttr {
// 	return &WidgetAttr{"compound", compound}
// }

// func LabelAttrText(text string) *WidgetAttr {
// 	return &WidgetAttr{"text", text}
// }

// func LabelAttrWidth(width int) *WidgetAttr {
// 	return &WidgetAttr{"width", width}
// }

// func LabelAttrPadding(padding Pad) *WidgetAttr {
// 	return &WidgetAttr{"padding", padding}
// }

// func LabelAttrState(state State) *WidgetAttr {
// 	return &WidgetAttr{"state", state}
// }

// func LabelAttrTakeFocus(takefocus bool) *WidgetAttr {
// 	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
// }

package label

import (
	"github.com/sysdeep/atk/tk"
)

// tk::label
type Label struct {
	tk.BaseWidget
	// common.Anchor
}

/*

идея - разложить всё по полочкам


*/

func NewLabel(parent tk.Widget, text string, attributes ...*tk.WidgetAttr) *Label {

	// вычисление какой вариант использовать tk or ttk
	// theme := checkInitUseTheme(attributes)

	iid := tk.MakeNamedWidgetId(parent, "tk_label")

	attributes = append(attributes, tk.NewWidgetAttr("text", text))

	info := tk.CreateWidgetInfo(iid, tk.WidgetTypeLabel, false, attributes)
	if info == nil {
		return nil
	}

	w := &Label{
		tk.NewBaseWidget(iid, info),
		// common.Anchor{},
	}

	tk.RegisterWidget(w)

	// NOTE: пример с миксинами, не получается сделать сцепку.. вложенная модель не видит внешнюю..
	// пример из tkd
	// w.SetTextAnchor("example")
	// w.ConfigureAnchor("aaa")

	return w
}

// func (w *Label) Attach(id string) error {
// 	info, err := tk.CheckWidgetInfo(id, tk.WidgetTypeLabel)
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
// 	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.Id()))
// }

// func (w *Label) Background() string {
// 	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.Id()))
// 	return r
// }

// func (w *Label) SetBorderWidth(width int) error {
// 	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.Id(), width))
// }

// func (w *Label) BorderWidth() int {
// 	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.Id()))
// 	return r
// }

// func (w *Label) SetForground(color string) error {
// 	setObjText("atk_tmp_text", color)
// 	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.Id()))
// }

// func (w *Label) Forground() string {
// 	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.Id()))
// 	return r
// }

// func (w *Label) SetReliefStyle(relief ReliefStyle) error {
// 	return eval(fmt.Sprintf("%v configure -relief {%v}", w.Id(), relief))
// }

// func (w *Label) ReliefStyle() ReliefStyle {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.Id()))
// 	return parserReliefStyleResult(r, err)
// }

// func (w *Label) SetFont(font Font) error {
// 	if font == nil {
// 		return ErrInvalid
// 	}
// 	return eval(fmt.Sprintf("%v configure -font {%v}", w.Id(), font.Id()))
// }

// func (w *Label) Font() Font {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.Id()))
// 	return parserFontResult(r, err)
// }

// func (w *Label) SetAnchor(anchor Anchor) error {
// 	return eval(fmt.Sprintf("%v configure -anchor {%v}", w.Id(), anchor))
// }

// func (w *Label) Anchor() Anchor {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -anchor", w.Id()))
// 	return parserAnchorResult(r, err)
// }

// func (w *Label) SetJustify(justify Justify) error {
// 	return eval(fmt.Sprintf("%v configure -justify {%v}", w.Id(), justify))
// }

// func (w *Label) Justify() Justify {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -justify", w.Id()))
// 	return parserJustifyResult(r, err)
// }

// func (w *Label) SetWrapLength(wraplength int) error {
// 	return eval(fmt.Sprintf("%v configure -wraplength {%v}", w.Id(), wraplength))
// }

// func (w *Label) WrapLength() int {
// 	r, _ := evalAsInt(fmt.Sprintf("%v cget -wraplength", w.Id()))
// 	return r
// }

// func (w *Label) SetImage(image *Image) error {
// 	if image == nil {
// 		return ErrInvalid
// 	}
// 	return eval(fmt.Sprintf("%v configure -image {%v}", w.Id(), image.Id()))
// }

// func (w *Label) Image() *Image {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -image", w.Id()))
// 	return parserImageResult(r, err)
// }

// func (w *Label) SetCompound(compound Compound) error {
// 	return eval(fmt.Sprintf("%v configure -compound {%v}", w.Id(), compound))
// }

// func (w *Label) Compound() Compound {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -compound", w.Id()))
// 	return parserCompoundResult(r, err)
// }

// func (w *Label) SetText(text string) error {
// 	setObjText("atk_tmp_text", text)
// 	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.Id()))
// }

// func (w *Label) Text() string {
// 	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.Id()))
// 	return r
// }

// func (w *Label) SetWidth(width int) error {
// 	return eval(fmt.Sprintf("%v configure -width {%v}", w.Id(), width))
// }

// func (w *Label) Width() int {
// 	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.Id()))
// 	return r
// }

// func (w *Label) SetPaddingN(padx int, pady int) error {
// 	if w.info.IsTtk {
// 		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.Id(), padx, pady))
// 	}
// 	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.Id(), padx, pady))
// }

// func (w *Label) PaddingN() (int, int) {
// 	var r string
// 	var err error
// 	if w.info.IsTtk {
// 		r, err = evalAsString(fmt.Sprintf("%v cget -padding", w.Id()))
// 	} else {
// 		r1, _ := evalAsString(fmt.Sprintf("%v cget -padx", w.Id()))
// 		r2, _ := evalAsString(fmt.Sprintf("%v cget -pady", w.Id()))
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
// 	return eval(fmt.Sprintf("%v configure -state {%v}", w.Id(), state))
// }

// func (w *Label) State() State {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.Id()))
// 	return parserStateResult(r, err)
// }

// func (w *Label) SetTakeFocus(takefocus bool) error {
// 	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.Id(), boolToInt(takefocus)))
// }

// func (w *Label) IsTakeFocus() bool {
// 	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.Id()))
// 	return r
// }

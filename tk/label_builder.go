package tk

import "fmt"

// experiments ----------------------------------------------------------------
/*
передавать в конструктор тонну опций муторно, нужен механизм который
1 - уменьшил бы кол-во набора
2 - гарантировал только поддерживаемый набор

// make
tk.NewLabelOptions().Text("test").
	Background("green").
	PadY(32).
	PadX(44).Foreground("red").
	Make(w).Pack()

// configure
lbl2 := tk.NewLabelOptions().Text("test configure").Make(w)
lbl2.Pack()
lbl2.Configure(tk.NewLabelOptions().Background("blue").GetOptions()...)

lbl3 := tk.NewLabelOptions().Text("test configure2").Make(w)
lbl3.Pack()
tk.NewLabelOptions().Background("black").Foreground("white").Configure(lbl3)

минусы:
- для конфигурирования(configure) всё равно остаётся больной список параметров,
необходимо придумать универсальный контейнер и для инициализации и для конфигурирования

*/

type LabelBuilder struct {
	options []OptionAdapter
	name    string
}

func NewLabelOptions() *LabelBuilder {
	return &LabelBuilder{
		options: []OptionAdapter{},
	}
}

func (o *LabelBuilder) PadX(value int) *LabelBuilder {
	o.options = append(o.options, OptPadX(value))
	return o
}

func (o *LabelBuilder) PadY(value int) *LabelBuilder {
	o.options = append(o.options, OptPadY(value))
	return o
}

func (o *LabelBuilder) Background(value string) *LabelBuilder {
	o.options = append(o.options, OptBackground(value))
	return o
}

func (o *LabelBuilder) ActiveBackground(value string) *LabelBuilder {
	o.options = append(o.options, OptActiveBackground(value))
	return o
}

func (o *LabelBuilder) Foreground(value string) *LabelBuilder {
	o.options = append(o.options, OptForeground(value))
	return o
}

func (o *LabelBuilder) ActiveForeground(value string) *LabelBuilder {
	o.options = append(o.options, OptActiveForeground(value))
	return o
}

func (o *LabelBuilder) Relief(value string) *LabelBuilder {
	o.options = append(o.options, OptRelief(value))
	return o
}

func (o *LabelBuilder) BorderWidth(value int) *LabelBuilder {
	o.options = append(o.options, OptBorderWidth(value))
	return o
}

func (o *LabelBuilder) Text(value string) *LabelBuilder {
	o.name = value
	o.options = append(o.options, OptText(value))
	return o
}

func (o *LabelBuilder) Justify(value string) *LabelBuilder {
	o.options = append(o.options, OptJustify(value))
	return o
}

func (o *LabelBuilder) Anchor(value string) *LabelBuilder {
	o.options = append(o.options, OptAnchor(value))
	return o
}

func (o *LabelBuilder) WrapLength(value int) *LabelBuilder {
	o.options = append(o.options, OptWrapLength(value))
	return o
}

func (o *LabelBuilder) Underline(value int) *LabelBuilder {
	o.options = append(o.options, OptUnderline(value))
	return o
}

func (o *LabelBuilder) Font(value string) *LabelBuilder {
	o.options = append(o.options, OptFont(value))
	return o
}

func (o *LabelBuilder) Compound(value string) *LabelBuilder {
	o.options = append(o.options, OptCompound(value))
	return o
}

func (o *LabelBuilder) Image(value *Image) *LabelBuilder {
	o.options = append(o.options, OptImage(value))
	return o
}

func (o *LabelBuilder) TakeFocus(value bool) *LabelBuilder {
	o.options = append(o.options, OptTakeFocus(value))
	return o
}

func (o *LabelBuilder) Make(parent Widget) *Label {
	return NewLabel(parent, o.name, o.options...)
}

func (o *LabelBuilder) GetOptions() []OptionAdapter {
	return o.options
}

func (o *LabelBuilder) Configure(w *Label) error {
	w.Configure(o.GetOptions()...)
	return nil
}

// func (o *LabelBuilder) CGet(w *Label) []OptionAdapter {
// 	for _, o := range o.options {
// 		w.CGet(o)
// 	}
// 	return o.options
// }

func ExampleBuilder() {

	b := NewLabelOptions().Text("My label").
		PadX(12).
		PadY(22).
		Background("red").Make(nil)

	fmt.Println(b)
}

package label

import (
	"fmt"

	"github.com/sysdeep/atk/tk"
)

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
	options []tk.OptionAdapter
	name    string
}

func NewLabelBuilder() *LabelBuilder {
	return &LabelBuilder{
		options: []tk.OptionAdapter{},
	}
}

func (o *LabelBuilder) PadX(value int) *LabelBuilder {
	o.options = append(o.options, PadX(value))
	return o
}

func (o *LabelBuilder) PadY(value int) *LabelBuilder {
	o.options = append(o.options, PadY(value))
	return o
}

func (o *LabelBuilder) Background(value string) *LabelBuilder {
	o.options = append(o.options, Background(value))
	return o
}

func (o *LabelBuilder) ActiveBackground(value string) *LabelBuilder {
	o.options = append(o.options, ActiveBackground(value))
	return o
}

func (o *LabelBuilder) Foreground(value string) *LabelBuilder {
	o.options = append(o.options, Foreground(value))
	return o
}

func (o *LabelBuilder) ActiveForeground(value string) *LabelBuilder {
	o.options = append(o.options, ActiveForeground(value))
	return o
}

func (o *LabelBuilder) Relief(value string) *LabelBuilder {
	o.options = append(o.options, Relief(value))
	return o
}

func (o *LabelBuilder) BorderWidth(value int) *LabelBuilder {
	o.options = append(o.options, BorderWidth(value))
	return o
}

func (o *LabelBuilder) Text(value string) *LabelBuilder {
	o.name = value
	o.options = append(o.options, Text(value))
	return o
}

func (o *LabelBuilder) Justify(value string) *LabelBuilder {
	o.options = append(o.options, Justify(value))
	return o
}

func (o *LabelBuilder) Anchor(value string) *LabelBuilder {
	o.options = append(o.options, Anchor(value))
	return o
}

func (o *LabelBuilder) WrapLength(value int) *LabelBuilder {
	o.options = append(o.options, WrapLength(value))
	return o
}

func (o *LabelBuilder) Underline(value int) *LabelBuilder {
	o.options = append(o.options, Underline(value))
	return o
}

func (o *LabelBuilder) Font(value string) *LabelBuilder {
	o.options = append(o.options, Font(value))
	return o
}

func (o *LabelBuilder) Compound(value string) *LabelBuilder {
	o.options = append(o.options, Compound(value))
	return o
}

func (o *LabelBuilder) Image(value *tk.Image) *LabelBuilder {
	o.options = append(o.options, Image(value))
	return o
}

func (o *LabelBuilder) TakeFocus(value bool) *LabelBuilder {
	o.options = append(o.options, TakeFocus(value))
	return o
}

func (o *LabelBuilder) Build(parent tk.Widget) *Label {
	return NewLabel(parent, o.name, o.options...)
}

// func (o *LabelBuilder) GetOptions() []OptionAdapter {
// 	return o.options
// }

func (o *LabelBuilder) Configure(w *Label) error {
	w.Configure(o.options...)
	return nil
}

// func (o *LabelBuilder) CGet(w *Label) []OptionAdapter {
// 	for _, o := range o.options {
// 		w.CGet(o)
// 	}
// 	return o.options
// }

func ExampleBuilder() {

	b := NewLabelBuilder().Text("My label").
		PadX(12).
		PadY(22).
		Background("red").Build(nil)

	fmt.Println(b)
}

package pack

import (
	"strconv"

	"github.com/sysdeep/atk/tk"
)

type PackBuilder struct {
	attrs map[string]string
}

func NewBuilder() *PackBuilder {
	return &PackBuilder{
		attrs: make(map[string]string),
	}
}

func (p *PackBuilder) SideTop() *PackBuilder {
	p.attrs["side"] = "top"
	return p
}

func (p *PackBuilder) SideLeft() *PackBuilder {
	p.attrs["side"] = "left"
	return p
}

func (p *PackBuilder) SideRight() *PackBuilder {
	p.attrs["side"] = "right"
	return p
}

func (p *PackBuilder) Anchor(anchor string) *PackBuilder {
	p.attrs["anchor"] = anchor
	return p
}

func (p *PackBuilder) PadX(v int) *PackBuilder {
	p.attrs["padx"] = strconv.Itoa(v)
	return p
}

func (p *PackBuilder) PadY(v int) *PackBuilder {
	p.attrs["pady"] = strconv.Itoa(v)
	return p
}

func (p *PackBuilder) FillBoth() *PackBuilder {
	p.attrs["fill"] = "both"
	return p
}

func (p *PackBuilder) Expand(b bool) *PackBuilder {
	v := "no"
	if b {
		v = "yes"
	}
	p.attrs["expand"] = v
	return p
}

func (p *PackBuilder) Build() []*tk.LayoutAttr {
	result := []*tk.LayoutAttr{}
	for k, v := range p.attrs {
		result = append(result, tk.NewLayoutAttr(k, v))
	}
	return result
}

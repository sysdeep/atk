package label

/*

set w .label
catch {destroy $w}
toplevel $w
wm title $w "Label Demonstration"
wm iconname $w "label"
positionWindow $w

label $w.msg -font $font -wraplength 4i -justify left -text "Five labels are displayed below: three textual ones on the left, and an image label and a text label on the right.  Labels are pretty boring because you can't do anything with them."
pack $w.msg -side top

## See Code / Dismiss buttons
set btns [addSeeDismiss $w.buttons $w]
pack $btns -side bottom -fill x

frame $w.left
frame $w.right
pack $w.left $w.right -side left -expand yes -padx 10 -pady 10 -fill both

label $w.left.l1 -text "First label"
label $w.left.l2 -text "Second label, raised" -relief raised
label $w.left.l3 -text "Third label, sunken" -relief sunken
pack $w.left.l1 $w.left.l2 $w.left.l3 -side top -expand yes -pady 2 -anchor w

# Main widget program sets variable tk_demoDirectory
image create photo label.ousterhout \
    -file [file join $tk_demoDirectory images ouster.png]
label $w.right.picture -borderwidth 2 -relief sunken -image label.ousterhout
label $w.right.caption -text "Tcl/Tk Creator"
pack $w.right.picture $w.right.caption -side top
*/

import (
	"github.com/sysdeep/atk/mtk/geo/pack"
	"github.com/sysdeep/atk/mtk/widgets/label"
	"github.com/sysdeep/atk/tk"
	"github.com/sysdeep/atk/tk/enum/anchor"
)

func NewLabel(parent tk.Widget) *tk.Frame {
	f := tk.NewFrame(parent)

	// intro
	intro := `Five labels are displayed below: three textual ones on the left, and an image label and a text label on the right.  Labels are pretty boring because you can't do anything with them.`
	// TODO: wrap inches
	tk.NewLabel(f, intro, tk.LabelOptWrapLength(400), tk.LabelOptJustify("left")).Pack(tk.PackAttrSideTop())

	// left
	l := tk.NewFrame(f)
	l.Pack(tk.PackAttrSideLeft(), tk.PackAttrExpand(true), tk.PackAttrPadx(10), tk.PackAttrPady(10), tk.PackAttrFillBoth())

	labelPackAttrs := []*tk.LayoutAttr{tk.PackAttrSideTop(), tk.PackAttrExpand(true), tk.PackAttrPady(2), tk.PackAttrAnchor(tk.AnchorWest)}
	tk.NewLabel(l, "First label").Pack(labelPackAttrs...)
	tk.NewLabel(l, "Second label, raised", label.Relief("raised")).Pack(labelPackAttrs...)
	tk.NewLabel(l, "Third label, sunken", label.Relief("sunken")).Pack(labelPackAttrs...)

	// right
	r := tk.NewFrame(f)
	r.Pack(tk.PackAttrSideLeft(), tk.PackAttrExpand(true), tk.PackAttrPadx(10), tk.PackAttrPady(10), tk.PackAttrFillBoth())

	// TODO: image
	tk.NewLabel(l, "", label.BorderWidth(2), label.Relief("sunken")).Pack(tk.PackAttrSideTop())
	tk.NewLabel(l, "Tcl/Tk Creator").Pack(tk.PackAttrSideTop())

	return f
}

const LabelTitle = "Label Demonstration"

func NewLabelWithBuilder(parent tk.Widget) *tk.Frame {
	f := tk.NewFrame(parent)

	intro := `Five labels are displayed below: three textual ones on the left, and an image label and a text label on the right.  Labels are pretty boring because you can't do anything with them.`
	// TODO: wrap inches
	// tk.NewLabel(f, intro, tk.LabelOptWrapLength(344), tk.LabelOptJustify("left")).Pack(tk.PackAttrSideTop())
	label.NewLabelBuilder().WrapLength(400).Justify("left").Text(intro).Build(f).Pack(pack.NewBuilder().SideTop().Build()...)

	l := tk.NewFrame(f)
	l.Pack(pack.NewBuilder().SideLeft().Expand(true).PadX(10).PadY(10).FillBoth().Build()...)

	labelPackAttrs := pack.NewBuilder().SideTop().Expand(true).PadY(2).Anchor(anchor.West).Build()
	label.NewLabel(l, "First label").Pack(labelPackAttrs...)
	label.NewLabel(l, "Second label, raised", label.Relief("raised")).Pack(labelPackAttrs...)
	label.NewLabel(l, "Third label, sunken", label.Relief("sunken")).Pack(labelPackAttrs...)

	r := tk.NewFrame(f)
	r.Pack(tk.PackAttrSideLeft(), tk.PackAttrExpand(true), tk.PackAttrPadx(10), tk.PackAttrPady(10), tk.PackAttrFillBoth())

	// TODO: image
	label.NewLabel(l, "", label.BorderWidth(2), label.Relief("sunken")).Pack(tk.PackAttrSideTop())
	label.NewLabel(l, "Tcl/Tk Creator").Pack(tk.PackAttrSideTop())

	return f
}

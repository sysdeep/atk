package main

import "github.com/sysdeep/atk/tk"

func doLabels1() {
	w := tk.NewWindow()

	intro := `Five labels are displayed below: three textual ones on the left, and an image label and a text label on the right.  Labels are pretty boring because you can't do anything with them.`
	tk.NewLabel(w, intro, tk.LabelOptWrapLength(344), tk.LabelOptJustify("left")).Pack(tk.PackAttrSideTop())

	// bottom
	controlsFrame := tk.NewFrame(w)
	tk.Pack(controlsFrame, tk.PackAttrSideBottom(), tk.PackAttrFillX())

	exitBtn := tk.NewButton(controlsFrame, "Exit")
	exitBtn.OnCommand(func() {
		w.Destroy()
	})
	tk.Pack(exitBtn, tk.PackAttrSideRight())

	// bind
	w.BindEvent("<Control-q>", func(e *tk.Event) {
		w.Destroy()
	})

	w.SetTitle("Label Demonstration")
	w.ResizeN(800, 600)

	// TODO
	// wm iconname $w "label"
	w.ShowNormal()
}

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

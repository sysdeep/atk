package main

import (
	"fmt"
	"log"

	"github.com/sysdeep/atk/tk"
)

/*

#--- text
text .n.widgetPane.lf.text -width 0 -height 6
.n.widgetPane.lf.text insert 1.0 "import std.stdio;\n\nvoid main(string[] args)\n{\n\twriteln(\"Hello World!\");\n}"
pack .n.widgetPane.lf.text -side bottom -fill both -anchor nw -padx 5

#--- entry
entry .n.widgetPane.lf.entry -textvariable myclock
pack .n.widgetPane.lf.entry -side left -fill x -anchor nw -padx 5 -expand true

#--- entry time ticker
proc clock:set var {
    global $var
    set $var [clock format [clock seconds] -format %H:%M:%S]
    after 1000 [list clock:set $var]
}

clock:set myclock          ;# call once, keeps ticking ;-) RS







#--- spinbox
ttk::spinbox .n.widgetPane.lf.spinbox -width 5 -wrap true -values "foo bar baz qux"
.n.widgetPane.lf.spinbox set foo
pack .n.widgetPane.lf.spinbox -side left -padx 5
			# auto entry3 = new SpinBox(entryLabelFrame)
			# 	.setData(["$foo", "[bar]", "\"baz\"", "{qux}"])
			# 	.setWrap(true)
			# 	.setWidth(5)
			# 	.pack(5, 0, GeometrySide.left);


#--- combobox
pack [ttk::combobox .n.widgetPane.lf.combobox -values "Option_1 Option_2 Option_3"] -side left -fill x -anchor nw -padx 5
.n.widgetPane.lf.combobox current 0

*/

func newWidgets(parent tk.Widget) *tk.Frame {
	f := tk.NewFrame(parent)
	f.SetReliefStyle(tk.ReliefStyleGroove)

	//--- text

	// 	 .n.widgetPane.lf.text -width 0 -height 6
	// .n.widgetPane.lf.text insert 1.0 "import std.stdio;\n\nvoid main(string[] args)\n{\n\twriteln(\"Hello World!\");\n}"
	// pack .n.widgetPane.lf.text -side bottom -fill both -anchor nw -padx 5

	// 	#--- entry
	// entry .n.widgetPane.lf.entry -textvariable myclock
	// pack .n.widgetPane.lf.entry -side left -fill x -anchor nw -padx 5 -expand true

	// #--- entry time ticker
	// proc clock:set var {
	//     global $var
	//     set $var [clock format [clock seconds] -format %H:%M:%S]
	//     after 1000 [list clock:set $var]
	// }

	// clock:set myclock          ;# call once, keeps ticking ;-) RS

	// #--- spinbox
	// ttk::spinbox .n.widgetPane.lf.spinbox -width 5 -wrap true -values "foo bar baz qux"
	// .n.widgetPane.lf.spinbox set foo
	// pack .n.widgetPane.lf.spinbox -side left -padx 5
	// 			# auto entry3 = new SpinBox(entryLabelFrame)
	// 			# 	.setData(["$foo", "[bar]", "\"baz\"", "{qux}"])
	// 			# 	.setWrap(true)
	// 			# 	.setWidth(5)
	// 			# 	.pack(5, 0, GeometrySide.left);

	// tk.NewComboBox(lf, tk.WidgetAttrInitUseTheme(true))
	// #--- combobox
	// pack [ttk::combobox .n.widgetPane.lf.combobox -values "Option_1 Option_2 Option_3"] -side left -fill x -anchor nw -padx 5
	// .n.widgetPane.lf.combobox current 0

	// buttons frame
	buttons_frame := tk.NewFrame(f)
	buttons_layout := tk.NewHPackLayout(buttons_frame)
	buttons_layout.AddWidget(make_buttons_frame(buttons_frame),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)
	buttons_layout.AddWidget(make_check_buttons_frame(buttons_frame),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)
	buttons_layout.AddWidget(make_radio_buttons_frame(buttons_frame),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)

	layout := tk.NewVPackLayout(f)
	layout.AddWidget(make_text_entry(f),
		tk.PackAttrFillX(),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
	)

	// pack .n.widgetPane.buttonsFrame -padx 10 -pady 10 -side left -fill both -expand true -anchor center
	layout.AddWidget(buttons_frame,
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PlaceAttrAnchor(tk.AnchorCenter),
	)
	// lll := tk.NewVPackLayout(f)

	// lll.AddWidgets(text, make_text_entry(f), entry, spin, combo)

	// pack [labelframe .n.widgetPane.lf -text "Text Entry"] -side top -fill both -padx 10 -pady 10

	return f
}

func make_text_entry(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent)
	main_frame.SetLabelText("Text Entry")

	// row 1
	row1 := tk.NewFrame(main_frame)

	// entry
	entry := tk.NewEntry(row1)

	// spinner
	spin := tk.NewSpinBox(row1)
	spinValues := []string{"foo", "bar", "baz", "qux"}
	log.Println(spinValues)
	spin.SetTextValues(spinValues)
	spin.Entry().SetWidth(5)

	// combo
	combo := tk.NewComboBox(row1, tk.WidgetAttrInitUseTheme(true))
	comboValues := []string{"option 1", "option 2", "option 3", "option 4"}
	combo.SetValues(comboValues)

	// row 1 layout
	pl := tk.NewHPackLayout(row1)
	pl.AddWidget(entry,
		tk.PackAttrSide(tk.SideLeft),
		tk.PackAttrFillX(),
		tk.PackAttrAnchor(tk.AnchorNorthWest),
		tk.PackAttrExpand(true))
	pl.AddWidget(spin,
		tk.PackAttrSide(tk.SideLeft),
		tk.PackAttrPadx(5))
	pl.AddWidget(combo,
		tk.PackAttrSide(tk.SideLeft),
		tk.PackAttrFillX(),
		tk.PackAttrAnchor(tk.AnchorNorthWest),
		tk.PackAttrPadx(5))

	// text
	text := tk.NewText(main_frame, tk.WidgetAttrHeight(6))
	text.InsertText(0, "import std.stdio;\n\nvoid main(string[] args)\n{\n\twriteln(\"Hello World!\");\n}")

	// main layout
	ml_pack := tk.NewVPackLayout(main_frame)
	ml_pack.AddWidget(row1,
		tk.PackAttrExpand(true),
		tk.PackAttrFillX(),
		tk.PackAttrPady(2),
		tk.PackAttrPadx(2),
	)
	ml_pack.AddWidget(text,
		tk.PackAttrFillX(),
		tk.PackAttrPady(2),
		tk.PackAttrPadx(2),
	)

	return main_frame
}

func make_buttons_frame(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent)
	main_frame.SetLabelText("Buttons")

	// text button
	text_button := tk.NewButton(main_frame, "Text button")

	// image button
	image, _ := tk.LoadImage("./assets/thumbnail.png")
	image_button := tk.NewButton(main_frame, "Image button", tk.ButtonAttrCompound(tk.CompoundLeft))
	image_button.SetImage(image)

	// menu button
	menu := tk.NewMenu(main_frame, tk.MenuAttrTearoff(false))
	for i := 1; i < 4; i++ {
		// NOTE: for go < 1.22 - need local variable for capture in closure
		local_i := i
		menu.AddAction(tk.NewActionEx(fmt.Sprintf("Option %d", i), func() {
			fmt.Println(fmt.Sprintf("Option %d was selected", local_i))
		}))
	}
	menu_button := tk.NewMenuButton(main_frame, "Menu button")
	menu_button.SetMenu(menu)

	main_layout := tk.NewVPackLayout(main_frame)
	main_layout.AddWidget(text_button, tk.PackAttrPadx(5), tk.PackAttrPady(5))
	main_layout.AddWidget(image_button, tk.PackAttrPadx(5), tk.PackAttrPady(5))
	main_layout.AddWidget(menu_button, tk.PackAttrPadx(5), tk.PackAttrPady(5))

	return main_frame
}

func make_check_buttons_frame(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent)
	main_frame.SetLabelText("Check buttons")

	main_layout := tk.NewVPackLayout(main_frame)

	// buttons
	var btn *tk.CheckButton
	is_selected := false
	for i := 1; i < 4; i++ {
		local_i := i
		btn = tk.NewCheckButton(main_frame, fmt.Sprintf("Option %d", local_i))
		main_layout.AddWidget(btn, tk.PackAttrPadx(5), tk.PackAttrPady(5))
		if !is_selected {
			btn.SetChecked(true)
			is_selected = true
		}
	}

	return main_frame
}

func make_radio_buttons_frame(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent)
	main_frame.SetLabelText("Radio buttons")

	main_layout := tk.NewVPackLayout(main_frame)

	// buttons
	group := tk.NewRadioGroup()
	is_selected := false
	var btn *tk.RadioButton
	for i := 1; i < 4; i++ {
		local_i := i
		btn = tk.NewRadioButton(main_frame, fmt.Sprintf("Option %d", local_i))
		main_layout.AddWidget(btn, tk.PackAttrPadx(5), tk.PackAttrPady(5))

		group.AddRadio(btn, i)

		if !is_selected {
			btn.SetChecked(true)
			is_selected = true
		}
	}

	return main_frame
}

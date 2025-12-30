package main

import (
	"fmt"
	"io/fs"
	"time"

	"github.com/sysdeep/atk/tk"
)

func makeWidgets(parent tk.Widget, assets fs.FS) *tk.Frame {
	f := tk.NewFrame(parent)

	// text entries
	tk.Pack(makeTextEntries(f),
		tk.PackAttrFillX(),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
	)

	// buttons frame
	buttons_frame := tk.NewFrame(f)

	tk.Pack(makeButtonsFrame(buttons_frame, assets),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)
	tk.Pack(makeCheckButtonsFrame(buttons_frame),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)
	tk.Pack(makeRadioButtonsFrame(buttons_frame),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrSideLeft(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)

	tk.Pack(buttons_frame,
		tk.PackAttrSideTop(),
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PlaceAttrAnchor(tk.AnchorCenter),
	)

	makeRangeScaleFrame(f)

	return f
}

func makeTextEntries(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Text Entry"))

	// row 1
	row1 := tk.NewFrame(main_frame)

	// entry
	entry := tk.NewEntry(row1)

	// spinner
	spin := tk.NewSpinBox(row1)
	spinValues := []string{"foo", "bar", "baz", "qux"}
	spin.SetTextValues(spinValues)
	spin.Entry().SetWidth(5)

	// combo
	combo := tk.NewComboBox(row1, tk.WidgetAttrInitUseTheme(true))
	comboValues := []string{"option 1", "option 2", "option 3", "option 4"}
	combo.SetValues(comboValues)
	combo.SetCurrentIndex(0)

	// row 1 layout
	tk.Pack(entry,
		tk.PackAttrSide(tk.SideLeft),
		tk.PackAttrFillX(),
		tk.PackAttrAnchor(tk.AnchorNorthWest),
		tk.PackAttrExpand(true))
	tk.Pack(spin,
		tk.PackAttrSide(tk.SideLeft),
		tk.PackAttrPadx(5))
	tk.Pack(combo,
		tk.PackAttrSide(tk.SideLeft),
		tk.PackAttrFillX(),
		tk.PackAttrAnchor(tk.AnchorNorthWest),
		tk.PackAttrPadx(5))

	// text
	text := tk.NewText(main_frame, tk.WidgetAttrHeight(6))
	text.InsertText(0, "import std.stdio;\n\nvoid main(string[] args)\n{\n\twriteln(\"Hello World!\");\n}")

	// main layout
	tk.Pack(row1,
		tk.PackAttrExpand(true),
		tk.PackAttrFillX(),
		tk.PackAttrPady(2),
		tk.PackAttrPadx(2),
	)
	tk.Pack(text,
		tk.PackAttrFillX(),
		tk.PackAttrPady(2),
		tk.PackAttrPadx(2),
	)

	// run time update
	go func() {
		for {
			tk.Async(func() {
				entry.SetText(time.Now().Format(time.TimeOnly))
			})
			<-time.Tick(1 * time.Second)
		}
	}()
	return main_frame
}

func makeButtonsFrame(parent tk.Widget, assets fs.FS) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Buttons"))

	// text button ----------------------------------------
	text_button := tk.NewButton(main_frame, "Text button")

	// image button ---------------------------------------
	image, _ := tk.LoadImageFrom("assets/thumbnail.png", assets)
	image_button := tk.NewButton(main_frame, "Image button", tk.ButtonAttrCompound(tk.CompoundLeft))
	image_button.SetImage(image)

	// menu button ----------------------------------------
	menu_button := tk.NewMenuButton(main_frame, "Menu button")
	// NOTE: parent of menu must be button
	menu := tk.NewMenu(menu_button, tk.MenuAttrTearoff(false))
	menu_button.SetMenu(menu)

	for i := 1; i < 4; i++ {
		// 	// NOTE: for go < 1.22 - need local variable for capture in closure
		local_i := i
		menu.AddAction(tk.NewActionEx(fmt.Sprintf("Option %d", local_i), func() {
			fmt.Printf("Option %d was selected\n", local_i)
		}))
	}

	tk.Pack(text_button, tk.PackAttrPadx(5), tk.PackAttrPady(5))
	tk.Pack(image_button, tk.PackAttrPadx(5), tk.PackAttrPady(5))
	tk.Pack(menu_button, tk.PackAttrPadx(5), tk.PackAttrPady(5))

	return main_frame
}

func makeCheckButtonsFrame(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Check buttons"))

	// buttons
	var btn *tk.CheckButton
	is_selected := true
	for i := 1; i < 4; i++ {
		local_i := i
		btn = tk.NewCheckButton(main_frame, fmt.Sprintf("Option %d", local_i))
		tk.Pack(btn, tk.PackAttrPadx(5), tk.PackAttrPady(5))

		btn.SetChecked(is_selected)
		is_selected = !is_selected
	}

	return main_frame
}

func makeRadioButtonsFrame(parent tk.Widget) *tk.LabelFrame {

	// main frame
	main_frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Radio buttons"))

	// buttons
	group := tk.NewRadioGroup()

	var btn *tk.RadioButton
	for i := 1; i < 4; i++ {
		local_i := i
		btn = tk.NewRadioButton(main_frame, fmt.Sprintf("Option %d", local_i))
		tk.Pack(btn, tk.PackAttrPadx(5), tk.PackAttrPady(5))
		group.AddRadio(btn, i)
	}

	group.SetCheckedIndex(0)

	return main_frame
}

func makeRangeScaleFrame(parent tk.Widget) *tk.LabelFrame {
	main_frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Progress & Scale"))

	// bar ------------------------------------------------
	bar := tk.NewProgressBar(main_frame, tk.Orient(tk.Horizontal), tk.ProgressBarAttrValue(4.0), tk.ProgressBarAttrMaximum(10))
	tk.Pack(
		bar,
		tk.PackAttrPadx(5),
		tk.PackAttrPady(5),
		tk.PackAttrSideTop(),
		tk.PackAttrFillX(),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)

	// scale ----------------------------------------------
	scale := tk.NewScale(
		main_frame,
		tk.Orient(tk.Horizontal),
		tk.ScaleAttrFrom(10),
		tk.ScaleAttrTo(0),
		tk.ScaleAttrShowValue(false),
		tk.ScaleAttrResolution(0.01),
	)
	tk.Pack(
		scale,
		tk.PackAttrPadx(5),
		tk.PackAttrPady(5),
		tk.PackAttrSideTop(),
		tk.PackAttrFillX(),
		tk.PackAttrAnchor(tk.AnchorCenter),
	)
	scale.SetValue(4)
	scale.OnCommand(func() {
		bar.SetValue(scale.Value())
	})

	// pack
	tk.Pack(
		main_frame,
		tk.PackAttrSideBottom(),
		tk.PackAttrFillBoth(),
		tk.PackAttrAnchor(tk.AnchorCenter),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
	)

	return main_frame
}

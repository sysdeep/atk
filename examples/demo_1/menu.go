package main

import "github.com/sysdeep/atk/tk"

func makeMenu(parent tk.Widget) *tk.Menu {
	menubar := tk.NewMenu(parent)

	// fm.AddAction(tk.NewAction("foo"))

	menubar.AddSubMenu("File", makeFileMenu(menubar))

	// // file
	// fileMenu := makeFileMenu(menubar)
	// menubar.AddCascade(tk.Lbl("File"), tk.Underline(0), tk.Mnu(fileMenu))

	// // themes
	// themesMenu := makeThemesMenu(menubar)
	// menubar.AddCascade(tk.Lbl("Themes"), tk.Underline(0), tk.Mnu(themesMenu))

	// // help
	// helpMenu := makeHelpMenu(menubar)
	// menubar.AddCascade(tk.Lbl("Help"), tk.Underline(0), tk.Mnu(helpMenu))

	return menubar

}

func makeFileMenu(parent tk.Widget) *tk.Menu {

	fileMenu := tk.NewMenu(parent, tk.MenuAttrTearoff(false))
	// 	// fileMenu.AddCommand(tk.Lbl("New"), tk.Underline(0), tk.Accelerator("Ctrl+N"))
	// 	// fileMenu.AddCommand(Lbl("Open..."), Underline(0), Accelerator("Ctrl+O"), Command(func() { GetOpenFile() }))
	// 	// Bind(App, "<Control-o>", Command(func() { fileMenu.Invoke(1) }))
	// 	// fileMenu.AddCommand(Lbl("Save"), Underline(0), Accelerator("Ctrl+S"))
	// 	// fileMenu.AddCommand(Lbl("Save As..."), Underline(5))
	// 	// fileMenu.AddCommand(Lbl("Close"), Underline(0), Accelerator("Crtl+W"))
	// 	// fileMenu.AddSeparator()
	// 	// fileMenu.AddCommand(Lbl("Exit"), Underline(1), Accelerator("Ctrl+Q"), ExitHandler())
	// 	// Bind(App, "<Control-q>", Command(func() { fileMenu.Invoke(6) }))

	// file check boxes
	fileCheckBoxes := fileMenu.AddNewSubMenu("Checkbutton submenu", tk.MenuAttrTearoff(false))
	fileCheckBoxes.AddAction(tk.NewCheckAction("Check option 1"))
	fileCheckBoxes.AddAction(tk.NewCheckAction("Check option 2"))
	fileCheckBoxes.AddAction(tk.NewCheckAction("Check option 3"))

	// file radio boxes
	radioGroup := tk.NewActionGroup()
	fileRadioBoxes := fileMenu.AddNewSubMenu("Radiobutton submenu", tk.MenuAttrTearoff(false))
	fileRadioBoxes.AddAction(radioGroup.AddNewRadioAction("Radio option 1"))
	fileRadioBoxes.AddAction(radioGroup.AddNewRadioAction("Radio option 2"))
	fileRadioBoxes.AddAction(radioGroup.AddNewRadioAction("Radio option 3"))
	radioGroup.SetCheckedIndex(1)

	fileMenu.AddSeparator()

	var exitAction = tk.NewActionEx("Exit", func() {
		tk.Quit()
	})
	// exitAction.SetData()
	fileMenu.AddAction(exitAction)

	// TODO:

	/*
		функционал не поддерживается
		.main_menu.file add command -label Exit -image img_close -compound left -accelerator "Ctrl-q" -command { exit }


		see AddAction
		script = fmt.Sprintf("%v add command -label $atk_tmp_label -command {%v}",
			w.id, act.actid)
	*/

	// fileMenu.AddCommand(
	// 	tk.Lbl("Exit"),
	// 	tk.Accelerator("Ctrl-q"),
	// 	tk.Underline(1),
	// 	tk.ExitHandler(),
	// 	tk.Image(tk.NewPhoto(tk.File("./media/cancel.png"))),
	// 	tk.Compound("left"),
	// )

	return fileMenu
}

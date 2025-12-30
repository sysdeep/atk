package main

import (
	"embed"

	"github.com/sysdeep/atk/tk"
)

//go:embed assets
var assetsFS embed.FS

type Window struct {
	*tk.Window
}

func NewWindow() *Window {

	tk.ExampleBuilder()

	// fmt.Println(assetsFS.ReadDir("assets"))

	tabWidgetsImage, _ := tk.LoadImageFrom("assets/layout_content.png", assetsFS)
	tabPanelsImage, _ := tk.LoadImageFrom("assets/application_tile_horizontal.png", assetsFS)
	tabCanvasImage, _ := tk.LoadImageFrom("assets/shape_ungroup.png", assetsFS)
	tabDialogsImage, _ := tk.LoadImageFrom("assets/application_double.png", assetsFS)

	mw := &Window{tk.RootWindow()}

	menu := makeMenu(mw)
	mw.SetMenu(menu)

	tabs := tk.NewNotebook(mw)

	tabWidgets := makeWidgets(tabs, assetsFS)
	tabPanels := newPanels(tabs)
	qlbl := tk.NewLabel(tabs, "Hello ATK")

	tabs.AddTab(tabWidgets, "Widgets", tk.TabAttrImage(tabWidgetsImage), tk.TabAttrCompound(tk.CompoundLeft))
	tabs.AddTab(tabPanels, "Panels", tk.TabAttrImage(tabPanelsImage), tk.TabAttrCompound(tk.CompoundLeft))
	tabs.AddTab(qlbl, "Canvas", tk.TabAttrImage(tabCanvasImage), tk.TabAttrCompound(tk.CompoundLeft))
	tabs.AddTab(tk.NewLabel(tabs, "Dialogs"), "Dialogs", tk.TabAttrImage(tabDialogsImage), tk.TabAttrCompound(tk.CompoundLeft))

	// layout
	tk.Pack(tabs,
		tk.PackAttrFillBoth(),
		tk.PackAttrPadx(10),
		tk.PackAttrPady(10),
		tk.PackAttrExpand(true),
	)

	// controls
	ctrlFrame := tk.NewFrame(mw)
	tk.Pack(ctrlFrame, tk.PackAttrSideBottom(), tk.PackAttrExpand(false), tk.PackAttrPadx(10), tk.PackAttrPady(10))

	btn := tk.NewButton(ctrlFrame, "Exit")
	btn.OnCommand(func() {
		tk.Quit()
	})
	tk.Pack(btn)

	mw.ResizeN(800, 600)
	mw.BindEvent("<Control-q>", func(e *tk.Event) {
		mw.Destroy()
	})
	return mw
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("ATK Sample")
		mw.Center(nil)
		mw.ShowNormal()
	})
}

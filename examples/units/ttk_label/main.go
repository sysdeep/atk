package main

import (
	"fmt"

	"github.com/sysdeep/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {

	// tk.SetMainTheme(nil)
	// fmt.Println("use theme?")
	// fmt.Println(tk.HasTheme())
	// fmt.Println(tk.MainTheme())

	w := tk.RootWindow()

	// labels frame
	labelsFrame := tk.NewFrame(w)
	tk.Pack(labelsFrame, tk.PackAttrSide(tk.SideTop), tk.PackAttrExpand(true), tk.PackAttrFillBoth())

	// make labels
	var labels = []*tk.TLabel{
		makeLabel1(labelsFrame),
		// makeLabel2(labelsFrame),
		// makeLabel3(labelsFrame),
		// makeLabelNN(labelsFrame),
		// makeLabelText(labelsFrame),
		// makeLabelWrapped(labelsFrame),
		// makeLabelFont(labelsFrame),
		// makeLabelBorderWidth(labelsFrame),
		// makeLabelUnderline(labelsFrame),
	}

	// grid labels
	for i, lbl := range labels {
		tk.Grid(lbl, tk.GridAttrRow(i), tk.GridAttrColumn(0))
	}

	framesGrid := tk.NewFrame(w)
	tk.Pack(framesGrid, tk.PackAttrSide(tk.SideTop), tk.PackAttrExpand(true), tk.PackAttrFillBoth())

	tk.Grid(makeFontsFrame(framesGrid), tk.GridAttrRow(0), tk.GridAttrColumn(0), tk.GridAttrIpadx(4), tk.GridAttrPadx(8))
	tk.Grid(makeReliefFrame(framesGrid), tk.GridAttrRow(0), tk.GridAttrColumn(1),
		tk.GridAttrIpadx(8), tk.GridAttrIpady(8), tk.GridAttrPadx(8))

	// ---------------------------------------------------
	// close
	btnClose := tk.NewButton(w, "Close")
	btnClose.OnCommand(func() {
		w.Destroy()
	})
	tk.Pack(btnClose, tk.PackAttrSide(tk.SideBottom))

	// bind
	w.BindEvent("<Control-q>", func(e *tk.Event) {
		w.Destroy()
	})
	return &Window{w}
}

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.SetTitle("Label Demo")
		w.Center(nil)
		w.ResizeN(800, 600)
		w.ShowNormal()
	})
}

func makeLabel1(parent tk.Widget) *tk.TLabel {
	// return tk.NewTLabel(parent, "Simple ttk label")
	return tk.NewTLabel2(parent, tk.NewTLabelOpts().Text("Simple ttk label")...)
}

func makeLabelUnderline(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Simple label with underline")
	lbl.Configure(tk.LabelOptUnderline(2))
	return lbl
}

func makeLabel2(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "BG and FG")
	lbl.Configure(
		tk.LabelOptBackground("green"),
		tk.LabelOptForeground("red"),
	)
	return lbl
}

func makeLabelFont(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Font")
	lbl.Configure(
		tk.LabelOptFont("TKHeadingFont"),
		// tk.LabelOptionForeground("red"),
	)
	return lbl
}

func makeLabel3(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Width = 30")
	lbl.Configure(
		tk.LabelOptBackground("green"),
		tk.LabelOptWidth(30),
		tk.LabelOptForeground("red"),
	)
	return lbl
}

func makeLabelBorderWidth(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "BorderWidth 3 and padx pady 10")
	lbl.Configure(
		tk.LabelOptBorderWidth(3),
		tk.LabelOptRelief("solid"),
		tk.LabelOptPadX(10),
		tk.LabelOptPadY(10),
	)
	return lbl
}

func makeLabelWrapped(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Wrap length 200")
	lbl.Configure(
		tk.LabelOptWrapLength(100),
	)
	return lbl
}

func makeLabelText(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "---")
	lbl.Configure(
		tk.LabelOptText("reconfigure text\ndisabled"),
		tk.LabelOptJustify("right"),
		tk.LabelOptWidth(30),
		tk.LabelOptBorderWidth(1),
		tk.LabelOptRelief("solid"),
		tk.LabelOptPadX(10),
		tk.LabelOptPadY(10),
		tk.LabelOptAnchor("e"),
		// tk.LabelOptionActiveBackground("green"),
		tk.LabelOptState("disabled"),
		tk.LabelOptHeight(13),
		// tk.LabelOptionBackground("blue"),
	)
	return lbl
}

func makeLabelNN(parent tk.Widget) *tk.Label {
	lbl_1 := tk.NewLabel(parent, "first")

	// test configure
	lbl_1.Configure(
		// tk.LabelOptionBackground("green"),
		// tk.LabelOptionForeground("red"),
		tk.LabelOptWidth(30),
		tk.OptReliefEnum(tk.ReliefStyleRidge),
	)

	// test try option throught cget
	bgq := tk.LabelOptBackground("")
	fmt.Printf("Bg before: %s\n", bgq.Value)
	lbl_1.CGet(bgq)
	fmt.Printf("Bg after: %s\n", bgq.Value)

	www := tk.LabelOptWidth(0)
	lbl_1.CGet(www)
	fmt.Printf("W after: %d\n", www.Value)

	rw := tk.OptReliefEnum(0)
	lbl_1.CGet(rw)
	fmt.Printf("relief: %s\n", rw.Value)

	return lbl_1
}

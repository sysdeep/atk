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
	return tk.NewTLabel(parent, "Simple ttk label")
}

func makeLabelUnderline(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Simple label with underline")
	lbl.Configure(tk.LabelOptionUnderline(2))
	return lbl
}

func makeLabel2(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "BG and FG")
	lbl.Configure(
		tk.LabelOptionBackground("green"),
		tk.LabelOptionForeground("red"),
	)
	return lbl
}

func makeLabelFont(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Font")
	lbl.Configure(
		tk.LabelOptionFont("TKHeadingFont"),
		// tk.LabelOptionForeground("red"),
	)
	return lbl
}

func makeLabel3(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Width = 30")
	lbl.Configure(
		tk.LabelOptionBackground("green"),
		tk.LabelOptionWidth(30),
		tk.LabelOptionForeground("red"),
	)
	return lbl
}

func makeLabelBorderWidth(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "BorderWidth 3 and padx pady 10")
	lbl.Configure(
		tk.LabelOptionBorderWidth(3),
		tk.LabelOptionRelief("solid"),
		tk.LabelOptionPadX(10),
		tk.LabelOptionPadY(10),
	)
	return lbl
}

func makeLabelWrapped(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Wrap length 200")
	lbl.Configure(
		tk.LabelOptionWrapLength(100),
	)
	return lbl
}

func makeLabelText(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "---")
	lbl.Configure(
		tk.LabelOptionText("reconfigure text\ndisabled"),
		tk.LabelOptionJustify("right"),
		tk.LabelOptionWidth(30),
		tk.LabelOptionBorderWidth(1),
		tk.LabelOptionRelief("solid"),
		tk.LabelOptionPadX(10),
		tk.LabelOptionPadY(10),
		tk.LabelOptionAnchor("e"),
		// tk.LabelOptionActiveBackground("green"),
		tk.LabelOptionState("disabled"),
		tk.LabelOptionHeight(13),
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
		tk.LabelOptionWidth(30),
		tk.OptionReliefEnum(tk.ReliefStyleRidge),
	)

	// test try option throught cget
	bgq := tk.LabelOptionBackground("")
	fmt.Printf("Bg before: %s\n", bgq.Value)
	lbl_1.CGet(bgq)
	fmt.Printf("Bg after: %s\n", bgq.Value)

	www := tk.LabelOptionWidth(0)
	lbl_1.CGet(www)
	fmt.Printf("W after: %d\n", www.Value)

	rw := tk.OptionReliefEnum(0)
	lbl_1.CGet(rw)
	fmt.Printf("relief: %s\n", rw.Value)

	return lbl_1
}

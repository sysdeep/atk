package main

import (
	"fmt"

	"github.com/sysdeep/atk/tk"
	"github.com/sysdeep/atk/tk/enum/cursor"
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

	framesGrid := tk.NewFrame(w)
	tk.Pack(framesGrid, tk.PackAttrSide(tk.SideTop), tk.PackAttrExpand(true), tk.PackAttrFillBoth())

	frames := [][]*tk.LabelFrame{
		{
			makeFontsFrame(framesGrid),
			makeReliefFrame(framesGrid),
			makeCursorFrame(framesGrid),
		},
		{
			makeOther(framesGrid),
			makeColorFrame(framesGrid),
		},
	}

	for i, row := range frames {
		for j, f := range row {
			tk.Grid(f, tk.GridAttrRow(i), tk.GridAttrColumn(j), tk.GridAttrPady(8), tk.GridAttrPadx(8))
		}
	}

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

func makeOther(parent tk.Widget) *tk.LabelFrame {
	frame := tk.NewLabelFrame(parent, tk.LabelFrameAttrLabelText("Other"))

	// make labels
	var labels = []*tk.Label{
		makeLabelSimple(frame),
		makeLabelWidth(frame),
		makeLabelNN(frame),
		makeLabelText(frame),
		makeLabelWrapped(frame),
		makeLabelFont(frame),
		makeLabelBorderWidth(frame),
		makeLabelUnderline(frame),
		makeLabelImage(frame),
	}

	// grid labels
	for i, lbl := range labels {
		tk.Grid(lbl, tk.GridAttrRow(i), tk.GridAttrColumn(0))
	}

	return frame
}

func makeLabelSimple(parent tk.Widget) *tk.Label {
	return tk.NewLabel(parent, "Simple label")
}

func makeLabelUnderline(parent tk.Widget) *tk.Label {
	lbl := tk.NewLabel(parent, "Simple label with underline")
	lbl.Configure(tk.LabelOptUnderline(2))
	return lbl
}

func makeLabelImage(parent tk.Widget) *tk.Label {
	img, _ := tk.LoadImage(tk.TkLibrary() + "/images/pwrdLogo100.gif")

	lbl := tk.NewLabel(parent, "with image",
		tk.LabelOptImage(img),
		tk.LabelOptCompound("left"),
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

func makeLabelWidth(parent tk.Widget) *tk.Label {
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
		tk.LabelOptCursor(cursor.Bogosity),
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

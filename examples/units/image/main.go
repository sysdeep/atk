package main

import (
	"image/color"
	"log"

	"github.com/sysdeep/atk/tk"
)

func main() {
	if err := tk.Init(); err != nil {
		log.Fatalln(err)
	}

	tk.MainLoop(func() {
		mw := tk.RootWindow()
		mw.SetTitle("Go Tk")

		img, err := tk.LoadImage("./liteide400.png", tk.ImageAttrTk85AlphaColor(color.Gray{0xef}))
		//img, err := tk.LoadImage(tk.TkLibrary() + "/images/pwrdLogo200.gif")
		if err != nil {
			log.Println(err)
		} else {
			lbl := tk.NewLabel(mw, "Image", tk.LabelOptCompound("center"))

			btn := tk.NewButton(mw, "Btn")
			btn.OnCommand(func() {
				lbl.Configure(tk.LabelOptImage(img))
			})
			tk.NewVPackLayout(mw).AddWidgets(btn, lbl)
			mw.SetSizeN(img.Size().Width, img.Size().Height+20)
		}
		mw.Center(nil)
		mw.ShowNormal()
	})
}

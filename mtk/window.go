package mtk

// Window represents a Tk window/widget. It implements common widget methods.
//
// Window implements Opt. When a Window instance is used as an Opt, it provides
// its path name.
type Window struct {
	fpath string
}

func (w *Window) isWidget() {}

// String implements fmt.Stringer.
func (w *Window) String() (r string) {
	if r = w.fpath; r == "" {
		r = "."
	}
	return r
}

func (w *Window) optionString(_ *Window) string {
	return w.String()
}

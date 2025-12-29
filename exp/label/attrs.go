package label

import "github.com/sysdeep/atk/tk"

func Background(color string) *tk.WidgetAttr {
	return tk.NewWidgetAttr("background", color)
}

func BorderWidth(width int) *tk.WidgetAttr {
	return tk.NewWidgetAttr("borderwidth", width)
}

func Foreground(color string) *tk.WidgetAttr {
	return tk.NewWidgetAttr("foreground", color)
}

func ReliefStyle(relief tk.ReliefStyle) *tk.WidgetAttr {
	return tk.NewWidgetAttr("relief", relief)
}

func Font(font tk.Font) *tk.WidgetAttr {
	if font == nil {
		return nil
	}
	return tk.NewWidgetAttr("font", font.Id())
}

func Anchor(anchor tk.Anchor) *tk.WidgetAttr {
	return tk.NewWidgetAttr("anchor", anchor)
}

func Justify(justify tk.Justify) *tk.WidgetAttr {
	return tk.NewWidgetAttr("justify", justify)
}

func WrapLength(wrapLength int) *tk.WidgetAttr {
	return tk.NewWidgetAttr("wraplength", wrapLength)
}

func Image(image *tk.Image) *tk.WidgetAttr {
	if image == nil {
		return nil
	}
	return tk.NewWidgetAttr("image", image.Id())
}

func Compound(compound tk.Compound) *tk.WidgetAttr {
	return tk.NewWidgetAttr("compound", compound)
}

func Text(text string) *tk.WidgetAttr {
	return tk.NewWidgetAttr("text", text)
}

func Width(width int) *tk.WidgetAttr {
	return tk.NewWidgetAttr("width", width)
}

func Padding(padding tk.Pad) *tk.WidgetAttr {
	return tk.NewWidgetAttr("padding", padding)
}

func State(state tk.State) *tk.WidgetAttr {
	return tk.NewWidgetAttr("state", state)
}

func TakeFocus(takeFocus bool) *tk.WidgetAttr {
	return tk.NewWidgetAttr("takefocus", boolToInt(takeFocus))
}

// utils ----------------------------------------------------------------------
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

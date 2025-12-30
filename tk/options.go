package tk

import (
	"fmt"
	"strconv"
)

// const (
// 	OptionBackground = "background"
// )

// type OptBackground struct{}

// func (b *OptBackground) String() string{
// 	return fmt.
// }

type OptionStringPair struct {
	Key   string
	Value string
}

type OptionAdapter interface {
	asStringPair() OptionStringPair
	asAnyPair(h func(string, interface{}) (interface{}, error)) error
}

// string ---------------------------------------------------------------------
type stringOption struct {
	Key   string
	Value string
}

func newStringOption(key string, value string) *stringOption {
	return &stringOption{
		Key:   key,
		Value: value,
	}
}

func (s stringOption) asStringPair() OptionStringPair {
	return OptionStringPair{
		Key:   s.Key,
		Value: s.Value,
	}
}

func (s *stringOption) asAnyPair(h func(k string, v interface{}) (interface{}, error)) error {
	raw, err := h(s.Key, s.Value)
	if err != nil {
		return err
	}

	r, ok := raw.(string)
	if ok {
		s.Value = r
		return nil
	}
	return fmt.Errorf("expected string, got %T", raw)
}

// int ---------------------------------------------------------------------
type intOption struct {
	Key   string
	Value int
}

func (s intOption) asStringPair() OptionStringPair {
	return OptionStringPair{
		Key:   s.Key,
		Value: strconv.Itoa(s.Value),
	}
}

func newIntOption(key string, value int) *intOption {
	return &intOption{
		Key:   key,
		Value: value,
	}
}

func (s *intOption) asAnyPair(h func(k string, v interface{}) (interface{}, error)) error {
	raw, err := h(s.Key, s.Value)
	if err != nil {
		return err
	}

	r, ok := raw.(int)
	if ok {
		s.Value = r
		return nil
	}
	return fmt.Errorf("expected int, got %T", raw)
}

// relief ---------------------------------------------------------------------
// type ReliefOption struct {
// 	Key   string
// 	Value string
// }

// func newReliefOption(key string, value ReliefStyle) *ReliefOption {
// 	return &ReliefOption{
// 		Key:   key,
// 		Value: value.String(),
// 	}
// }

// func (s ReliefOption) asStringPair() OptionStringPair {
// 	return OptionStringPair{
// 		Key:   s.Key,
// 		Value: s.Value,
// 	}
// }

// func (s *ReliefOption) asAnyPair(h func(k string, v interface{}) (interface{}, error)) error {
// 	raw, err := h(s.Key, s.Value)
// 	if err != nil {
// 		return err
// 	}

// 	r, ok := raw.(string)
// 	if ok {
// 		s.Value = s.parse(r)
// 		return nil
// 	}
// 	return fmt.Errorf("expected string, got %T", raw)
// }

// func (s *ReliefOption) parse(value string) ReliefStyle {
// 	for n, s := range borderStyleName {
// 		if s == value {
// 			return ReliefStyle(n)
// 		}
// 	}

// 	return ReliefStyle(-1)
// }

// options --------------------------------------------------------------------
/*
Command-Line Name: -background or -bg
Database Name: background
Database Class: Background
Specifies the normal background color to use when displaying the widget.
*/
func OptionBackground(color string) *stringOption {
	return newStringOption("background", color)
}

/*
Command-Line Name: -activebackground
Database Name: activeBackground
Database Class: Foreground
Specifies background color to use when drawing active elements. An element (a widget or portion of a widget) is active if the mouse cursor is positioned over the element and pressing a mouse button will cause some action to occur. If strict Motif compliance has been requested by setting the tk_strictMotif variable, this option will normally be ignored; the normal background color will be used instead. For some elements on Windows and Macintosh systems, the active color will only be used while mouse button 1 is pressed over the element.
*/
func OptionActiveBackground(color string) *stringOption {
	return newStringOption("activebackground", color)
}

/*
Command-Line Name: -foreground or -fg
Database Name: foreground
Database Class: Foreground
Specifies the normal foreground color to use when displaying the widget.
*/
func OptionForeground(color string) *stringOption {
	return newStringOption("foreground", color)
}

/*
Command-Line Name: -activeforeground
Database Name: activeForeground
Database Class: Background
Specifies foreground color to use when drawing active elements. See above for definition of active elements.
*/
func OptionActiveForeground(color string) *stringOption {
	return newStringOption("activeforeground", color)
}

/*
Command-Line Name: -borderwidth or -bd
Database Name: borderWidth
Database Class: BorderWidth
Specifies a non-negative value indicating the width of the 3-D border to draw around the outside of the widget (if such a border is being drawn; the -relief option typically determines this). The value may also be used when drawing 3-D effects in the interior of the widget. The value may have any of the forms acceptable to Tk_GetPixels.
*/
func OptionBorderWidth(width int) *intOption {
	return newIntOption("borderwidth", width)
}

/*
Command-Line Name: -padx
Database Name: padX
Database Class: Pad
Specifies a non-negative value indicating how much extra space to request for the widget in the X-direction. The value may have any of the forms acceptable to Tk_GetPixels. When computing how large a window it needs, the widget will add this amount to the width it would normally need (as determined by the width of the things displayed in the widget); if the geometry manager can satisfy this request, the widget will end up with extra internal space to the left and/or right of what it displays inside. Most widgets only use this option for padding text: if they are displaying a bitmap or image, then they usually ignore padding options.
*/
func OptionPadX(value int) *intOption {
	return newIntOption("padx", value)
}

/*
Command-Line Name: -pady
Database Name: padY
Database Class: Pad
Specifies a non-negative value indicating how much extra space to request for the widget in the Y-direction. The value may have any of the forms acceptable to Tk_GetPixels. When computing how large a window it needs, the widget will add this amount to the height it would normally need (as determined by the height of the things displayed in the widget); if the geometry manager can satisfy this request, the widget will end up with extra internal space above and/or below what it displays inside. Most widgets only use this option for padding text: if they are displaying a bitmap or image, then they usually ignore padding options.
*/
func OptionPadY(value int) *intOption {
	return newIntOption("pady", value)
}

// TODO: перевести ReliefStyle на строки-константы
func OptionReliefEnum(relief ReliefStyle) *stringOption {
	return newStringOption("relief", relief.String())
}

/*
Command-Line Name: -relief
Database Name: relief
Database Class: Relief
Specifies the 3-D effect desired for the widget. Acceptable values are raised, sunken, flat, ridge, solid, and groove. The value indicates how the interior of the widget should appear relative to its exterior; for example, raised means the interior of the widget should appear to protrude from the screen, relative to the exterior of the widget.
*/
func OptionRelief(relief string) *stringOption {
	return newStringOption("relief", relief)
}

/*
Command-Line Name: -text
Database Name: text
Database Class: Text
Specifies a string to be displayed inside the widget. The way in which the string is displayed depends on the particular widget and may be determined by other options, such as -anchor or -justify.
*/
func OptionText(value string) *stringOption {
	return newStringOption("text", value)
}

/*
Command-Line Name: -justify
Database Name: justify
Database Class: Justify
When there are multiple lines of text displayed in a widget, this option determines how the lines line up with each other. Must be one of left, center, or right. Left means that the lines' left edges all line up, center means that the lines' centers are aligned, and right means that the lines' right edges line up.
*/
func OptionJustify(value string) *stringOption {
	// TODO: justify variants
	return newStringOption("justify", value)
}

/*
Command-Line Name: -anchor
Database Name: anchor
Database Class: Anchor
Specifies how the information in a widget (e.g. text or a bitmap) is to be displayed in the widget. Must be one of the values n, ne, e, se, s, sw, w, nw, or center. For example, nw means display the information such that its top-left corner is at the top-left corner of the widget.
*/
func OptionAnchor(value string) *stringOption {
	// TODO: anchor variants
	return newStringOption("anchor", value)
}

/*
Command-Line Name: -wraplength
Database Name: wrapLength
Database Class: WrapLength
For widgets that can perform word-wrapping, this option specifies the maximum line length. Lines that would exceed this length are wrapped onto the next line, so that no line is longer than the specified length. The value may be specified in any of the standard forms for screen distances. If this value is less than or equal to 0 then no wrapping is done: lines will break only at newline characters in the text.
*/
func OptionWrapLength(value int) *intOption {
	return newIntOption("wraplength", value)
}

/*
Command-Line Name: -font
Database Name: font
Database Class: Font
Specifies the font to use when drawing text inside the widget. The value may have any of the forms described in the font manual page under FONT DESCRIPTION.
*/
func OptionFont(value string) *stringOption {
	// TODO: anchor variants
	return newStringOption("font", value)
}

/*
Command-Line Name: -underline
Database Name: underline
Database Class: Underline
Specifies the integer index of a character to underline in the widget. This option is used by the default bindings to implement keyboard traversal for menu buttons and menu entries. 0 corresponds to the first character of the text displayed in the widget, 1 to the next character, and so on.
*/
func OptionUnderline(value int) *intOption {
	return newIntOption("underline", value)
}

/*
Command-Line Name: -compound
Database Name: compound
Database Class: Compound
Specifies if the widget should display text and bitmaps/images at the same time, and if so,
where the bitmap/image should be placed relative to the text.
Must be one of the values none, bottom, top, left, right, or center.
For example, the (default) value none specifies that the bitmap or image should (if defined)
be displayed instead of the text, the value left specifies that the bitmap or
image should be displayed to the left of the text, and the value center specifies that the bitmap or image should be displayed on top of the text.
*/
func OptionCompound(value string) *stringOption {
	// TODO: variants
	return newStringOption("compound", value)
}

/*
Command-Line Name: -image
Database Name: image
Database Class: Image
Specifies an image to display in the widget, which must have been created with the image create command. Typically, if the -image option is specified then it overrides other options that specify a bitmap or textual value to display in the widget, though this is controlled by the -compound option; the -image option may be reset to an empty string to re-enable a bitmap or text display.
*/
func OptionImage(image *Image) *stringOption {
	return newStringOption("image", image.Id())
}

/*
Command-Line Name: -takefocus
Database Name: takeFocus
Database Class: TakeFocus
Determines whether the window accepts the focus during keyboard traversal (e.g., Tab and Shift-Tab). Before setting the focus to a window, the traversal scripts consult the value of the -takefocus option. A value of 0 means that the window should be skipped entirely during keyboard traversal. 1 means that the window should receive the input focus as long as it is viewable (it and all of its ancestors are mapped). An empty value for the option means that the traversal scripts make the decision about whether or not to focus on the window: the current algorithm is to skip the window if it is disabled, if it has no key bindings, or if it is not viewable. If the value has any other form, then the traversal scripts take the value, append the name of the window to it (with a separator space), and evaluate the resulting string as a Tcl script. The script must return 0, 1, or an empty string: a 0 or 1 value specifies whether the window will receive the input focus, and an empty string results in the default decision described above. Note: this interpretation of the option is defined entirely by the Tcl scripts that implement traversal: the widget implementations ignore the option entirely, so you can change its meaning if you redefine the keyboard traversal scripts.
*/
func OptionTakeFocus(value bool) *intOption {
	intValue := boolToInt(value)
	return newIntOption("takefocus", intValue)
}

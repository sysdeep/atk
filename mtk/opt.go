package mtk

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"strings"
	"time"
)

var (
	// https://pdos.csail.mit.edu/archive/rover/RoverDoc/escape_shell_table.html
	//
	// The following characters are dissallowed or have special meanings in Tcl and
	// so are escaped:
	//
	//	&;`'"|*?~<>^()[]{}$\
	badChars = [...]bool{
		' ':  true,
		'"':  true,
		'$':  true,
		'&':  true,
		'(':  true,
		')':  true,
		'*':  true,
		';':  true,
		'<':  true,
		'>':  true,
		'?':  true,
		'[':  true,
		'\'': true,
		'\\': true,
		'\n': true,
		'\r': true,
		'\t': true,
		']':  true,
		'^':  true,
		'`':  true,
		'{':  true,
		'|':  true,
		'}':  true,
		'~':  true,
	}
)

// func winCollect(w *Window, options ...Opt) string {
// 	var a []string
// 	for _, v := range options {
// 		a = append(a, v.optionString(w))
// 	}
// 	return strings.Join(a, " ")
// }

// func collectAny(options ...any) string {
// 	var a []string
// 	for _, v := range options {
// 		switch x := v.(type) {
// 		case Opt:
// 			a = append(a, x.optionString(nil))
// 		default:
// 			a = append(a, tclSafeString(fmt.Sprint(x)))
// 		}
// 	}
// 	return strings.Join(a, " ")
// }

// func collect(options ...Opt) string {
// 	var a []string
// 	for _, v := range options {
// 		a = append(a, v.optionString(nil))
// 	}
// 	return strings.Join(a, " ")
// }

// func collectOne(name string, options ...Opt) string {
// 	for _, v := range options {
// 		opt := v.optionString(nil)
// 		if strings.HasPrefix(opt, name) {
// 			return strings.TrimSpace(opt[len(name):])
// 		}
// 	}
// 	return ""
// }

// Opts is a list of options. It implements Opt.
// type Opts []Opt

// func (o Opts) optionString(w *Window) string {
// 	return winCollect(w, []Opt(o)...)
// }

// Opt represents an optional argument.
type Opt interface {
	optionString(w *Window) string
}

type rawOption string

func (s rawOption) optionString(w *Window) string {
	return string(s)
}

type stringOption string

func (s stringOption) optionString(w *Window) string {
	return tclSafeString(string(s))
}

// Txt option.
//
// Specifies a string to be displayed inside the widget.  The way in which
// the string is displayed depends on the particular widget and may be
// determined by other options, such as -anchor or -justify.
//
// Known uses:
//   - [Button]
//   - [Checkbutton]
//   - [Label]
//   - [Labelframe]
//   - [Menubutton]
//   - [Message]
//   - [Radiobutton]
//   - [TButton]
//   - [TCheckbutton]
//   - [TLabel]
//   - [TLabelframe] (widget specific)
//   - [TMenubutton]
//   - [TProgressbar]
//   - [TRadiobutton]
func Txt(val any) Opt {
	return rawOption(fmt.Sprintf(`-text %s`, optionString(val)))
}

// Returns a single Tcl string, no braces, except "{}" is returned for s == "".
func tclSafeString(s string) string {
	if s == "" {
		return "{}"
	}

	const badString = "&;`'\"|*?~<>^()[]{}$\\\n\r\t "
	if strings.ContainsAny(s, badString) {
		var b strings.Builder
		for _, c := range s {
			switch {
			case int(c) < len(badChars) && badChars[c]:
				fmt.Fprintf(&b, "\\x%02x", c)
			default:
				b.WriteRune(c)
			}
		}
		s = b.String()
	}
	return s
}

func hexDigit(b byte) byte {
	if b <= 9 {
		return '0' + b
	}

	return 'a' + b - 10
}

func tclBinaryString(s string) string {
	var b strings.Builder
	for _, v := range s {
		b.WriteString(fmt.Sprintf("\\U%06X", v))
	}
	return b.String()
}

func tclBinaryBytes(s []byte) string {
	var b strings.Builder
	for _, v := range s {
		b.WriteByte('\\')
		b.WriteByte('x')
		b.WriteByte(hexDigit(v >> 4))
		b.WriteByte(hexDigit(v & 15))
	}
	return b.String()
}

var xpmSig = []byte("/* XPM */") // https://en.wikipedia.org/wiki/X_PixMap

func optionString(v any) (r string) {
	switch x := v.(type) {
	case time.Duration:
		return fmt.Sprint(int64((x + time.Millisecond/2) / time.Millisecond))
	case []byte:
		switch {
		case bytes.HasPrefix(x, xpmSig):
			return tclSafeString(fmt.Sprintf("%s", v))
		default:
			return tclBinaryBytes(x)
		}
	case image.Image:
		var buf bytes.Buffer
		err := png.Encode(&buf, x)
		if err != nil {
			fail(err)
			return ""
		}

		return base64.StdEncoding.EncodeToString(buf.Bytes())
	case []FileType:
		var a []string
		for _, v := range x {
			a = append(a, fmt.Sprintf("{%s {%s} %s}", tclSafeString(v.TypeName), tclSafeStrings(v.Extensions...), v.MacType))
		}
		return fmt.Sprintf("{%s}", strings.Join(a, " "))
	default:
		return tclSafeString(fmt.Sprint(v))
	}
}

// Returns a space separated list of safe Tcl strings.
func tclSafeStrings(s ...string) string {
	var a []string
	for _, v := range s {
		a = append(a, tclSafeString(v))
	}
	return strings.Join(a, " ")
}

// FileType specifies a single file type for the [Filetypes] option.
type FileType struct {
	TypeName   string   // Eg. "Go files"
	Extensions []string // Eg. []string{".go"}
	MacType    string   // Eg. "TEXT"
}

// ErrorMode selects the action taken on errors.
var ErrorMode int

// Error records errors when [CollectErrors] is true.
var Error error

// Error modes
const (
	// Errors will panic with a stack trace.
	PanicOnError = iota
	// Errors will be recorded into the Error variable using errors.Join
	CollectErrors
)

func fail(err error) {
	switch ErrorMode {
	default:
		fallthrough
	case PanicOnError:
		// TODO
		// if dmesgs {
		// 	dmesg("PANIC %v", err)
		// }
		panic(err)
	case CollectErrors:
		Error = errors.Join(Error, err)
	}
}

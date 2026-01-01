package state

type State string

// Command-Line Name: -state
// Database Name: state
// Database Class: State
// May be set to normal or disabled to control the disabled state bit. This is a write-only option: setting it changes the widget state, but the state widget command does not affect the -state option.
const (
	Normal   State = "normal"
	Active   State = "active"
	Disabled State = "disabled"
)

func (s State) String() string {
	return string(s)
}

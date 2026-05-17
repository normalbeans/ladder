package key

type Key string

// add more here ig
const (
	CtrlC      Key = "\x03"
	BackSpace  Key = "\x08"
	Escape     Key = "\x1b"
	ArrowUp    Key = "\x1b[A"
	ArrowDown  Key = "\x1b[B"
	ArrowRight Key = "\x1b[C"
	ArrowLeft  Key = "\x1b[D"
)

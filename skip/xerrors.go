// This file holds xerrors specific logic. xerrors was meant to be the
// behaviour that would land in the standard library in go 1.13. Not all
// of it did, but this still remains the default for vice. Still, it is kept
// in its own file to make it easier to remove or put behind a build tag.

package skip

import (
	"fmt"

	"golang.org/x/xerrors"
)

func (e *sealError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *sealError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

func (f frame) Format(p xerrors.Printer) {
	if !p.Detail() {
		return
	}

	fn, file, line := f.detail()
	if fn != "" {
		p.Printf("%s\n    ", fn)
	}
	if file != "" {
		p.Printf("%s:%d\n", file, line)
	}
}

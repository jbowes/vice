// frame reimplements stack frame handling in a way that is compatible with
// multiple error ecosystems.

package skip

import "runtime"

type frame [3]uintptr

func caller(skip int) frame {
	var f frame
	runtime.Callers(skip+1, f[:])
	return f
}

func (f frame) detail() (string, string, int) {
	frs := runtime.CallersFrames(f[:])

	if _, ok := frs.Next(); ok {
		if fr, ok := frs.Next(); ok {
			return fr.Function, fr.File, fr.Line
		}
	}

	return "", "", 0
}

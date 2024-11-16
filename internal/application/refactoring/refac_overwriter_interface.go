package refactoring

import "io"

type RefacOverWriter interface {
	OverWrite(w io.Writer, src string)
}

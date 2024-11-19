package diff

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type CmpDiffer struct{}

func NewCmpDiffer() *CmpDiffer {
	return &CmpDiffer{}
}

func (CmpDiffer) Diff(originSrc string, refactSrc string) string {
	diff := cmp.Diff(originSrc, refactSrc)
	return fmt.Sprintf("the difference is (-orign +refactored):%s", diff)
}

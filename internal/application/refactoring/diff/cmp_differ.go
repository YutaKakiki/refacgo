package diff

type CmpDiffer struct{}

func NewCmpDiffer() *CmpDiffer {
	return &CmpDiffer{}
}

func (CmpDiffer) Diff(originSrc string, refactSrc string) string {
	return ""
}

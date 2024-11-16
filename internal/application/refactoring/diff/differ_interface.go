package diff

type Differ interface {
	Diff(originSrc, refactSrc string) string
}

package presenter

type RefacOverWriter interface {
	OverWrite(filename string, src string) error
	OverWriteWithHeaderComment(filename string, src string) error
}

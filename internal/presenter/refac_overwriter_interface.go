package presenter

//go:generate mockgen -package=presenter -source=./refac_overwriter_interface.go -destination=./refac_overwriter_mock.go
type RefacOverWriter interface {
	OverWrite(filename string, src string) error
	OverWriteWithHeaderComment(filename string, src string) error
}

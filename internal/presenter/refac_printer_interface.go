package presenter

//go:generate mockgen -package=presenter -source=./refac_printer_interface.go -destination=./refac_printer_mock.go

type RefacPrinter interface {
	Print(text ...string)
}

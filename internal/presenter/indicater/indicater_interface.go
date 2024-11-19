package indicater

//go:generate mockgen -package=indicater -source=./indicater_interface.go -destination=./indicater_mock.go
type Indicater interface {
	Start()
	Stop()
}

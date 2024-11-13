package indicater

import (
	"time"

	"github.com/briandowns/spinner"
)

type Indicater struct {
	Spinner *spinner.Spinner
}

func NewIndicater() *Indicater {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	return &Indicater{
		Spinner: s,
	}
}

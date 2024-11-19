package indicater

import (
	"time"

	"github.com/briandowns/spinner"
)

func NewEvalSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " waiting for evaluating ......"
	return s
}

package indicater

import (
	"time"

	"github.com/briandowns/spinner"
)

func NewRefacSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " waiting for refactoring ......"
	return s
}

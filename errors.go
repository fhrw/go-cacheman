package cacheman

import (
	"errors"
)

// some errors
var ErrNotAvailable = errors.New("Requested value not available")
var ErrFarWriteFail = errors.New("Couldn't write that to the far")
var ErrNearWriteFail = errors.New("Couldn't write that to the near")

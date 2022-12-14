package definition

import (
	"errors"
)

var (
	ErrDataNotFound = errors.New("Data not found.")
	ErrBadRequest = errors.New("Bad request.")
	ErrInternalServer = errors.New("Can't process this request at the moment.")
)
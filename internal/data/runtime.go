package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshalJSON() ([]byte, error) {
	val := fmt.Sprintf("%d mins", r)

	return []byte(strconv.Quote(val)), nil
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	unquotedVal, err := strconv.Unquote(string(data))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedVal, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	val, err := strconv.ParseInt(parts[0], 10, 32)

	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(val)

	return nil
}

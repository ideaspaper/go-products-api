package service

import "errors"

var (
	ErrServiceUnknown      = errors.New("ERROR: [Service] Unknown")
	ErrServiceDataNotFound = errors.New("ERROR: [Service] Data Not Found")
)

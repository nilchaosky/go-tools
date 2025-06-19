package tools_redis

import "errors"

type NotClientError struct {
}

func (e *NotClientError) Error() string {
	return "redis client is nil "
}

func IsNotClient(err error) bool {
	if err == nil {
		return false
	}
	var e *NotClientError
	return errors.As(err, &e)
}

type NotFoundError struct {
	label string
}

func (e *NotFoundError) Error() string {
	return "redis key: " + e.label + " not found"
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

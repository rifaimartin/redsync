package redsync

import "errors"

// ErrFailed this variable new ErrFailed
var ErrFailed = errors.New("redsync: failed to acquire lock")

package bdstockexchange

import "errors"

var (
	// errInvalidGroupName is thrown when user input an illegal group name
	errInvalidGroupName = errors.New("group name is invalid. enter a valid group name. ex : A, B, G, N, Z")
)

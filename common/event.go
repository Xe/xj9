package common

import "github.com/matrix-org/gomatrix"

// Event is a wrapper tuple for an incoming event.
type Event struct {
	*gomatrix.Client
	*gomatrix.Event
}

package integration

import (
	"github.com/sergiolaverde0/mangal/integration/anilist"
	"github.com/sergiolaverde0/mangal/source"
)

// Integrator is the interface that wraps the basic integration methods.
type Integrator interface {
	// MarkRead marks a chapter as read
	MarkRead(chapter *source.Chapter) error
}

var (
	Anilist Integrator = anilist.New()
)

package packer

import (
	"fmt"
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/packer/cbz"
	"github.com/sergiolaverde0/mangal/packer/plain"
	"github.com/sergiolaverde0/mangal/packer/zip"
	"github.com/sergiolaverde0/mangal/source"
	"github.com/samber/lo"
	"strings"
)

// Packer is the interface that all packers must implement.
type Packer interface {
	// SupportedConversion what conversion format does this packer supports
	SupportedConversion() (formats []constant.ConversionFormat)
	Save(chapter *source.Chapter) (string, error)
	SaveTemp(chapter *source.Chapter) (string, error)
}

var packers = map[string]Packer{
	constant.FormatPlain: plain.New(),
	constant.FormatCBZ:   cbz.New(),
	constant.FormatZIP:   zip.New(),
}

// Available returns a list of available packers.
func Available() []string {
	return lo.Keys(packers)
}

// Get returns a packer by name.
// If the packer is not available, an error is returned.
func Get(name string) (Packer, error) {
	if packer, ok := packers[name]; ok {
		return packer, nil
	}

	return nil, fmt.Errorf("unkown format \"%s\", available options are %s", name, strings.Join(Available(), ", "))
}

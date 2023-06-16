package provider

import (
	"github.com/belphemur/mangal/provider/asurascans"
	"github.com/belphemur/mangal/provider/flamescans"
	"github.com/belphemur/mangal/provider/generic"
	"github.com/belphemur/mangal/provider/mangadex"
	"github.com/belphemur/mangal/provider/manganato"
	"github.com/belphemur/mangal/provider/manganelo"
	"github.com/belphemur/mangal/provider/mangapill"
	"github.com/belphemur/mangal/source"
)

const CustomProviderExtension = ".lua"

var builtinProviders = []*Provider{
	{
		ID:   mangadex.ID,
		Name: mangadex.Name,
		CreateSource: func() (source.Source, error) {
			return mangadex.New(), nil
		},
	},
}

func init() {
	for _, conf := range []*generic.Configuration{
		manganelo.Config,
		manganato.Config,
		mangapill.Config,
		flamescans.Config,
		asurascans.Config,
	} {
		conf := conf
		builtinProviders = append(builtinProviders, &Provider{
			ID:   conf.ID(),
			Name: conf.Name,
			CreateSource: func() (source.Source, error) {
				return generic.New(conf), nil
			},
		})
	}
}

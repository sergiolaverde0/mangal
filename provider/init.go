package provider

import (
	"github.com/sergiolaverde0/mangal/provider/asurascans"
	"github.com/sergiolaverde0/mangal/provider/flamescans"
	"github.com/sergiolaverde0/mangal/provider/generic"
	"github.com/sergiolaverde0/mangal/provider/mangadex"
	"github.com/sergiolaverde0/mangal/provider/manganato"
	"github.com/sergiolaverde0/mangal/provider/manganelo"
	"github.com/sergiolaverde0/mangal/provider/mangapill"
	"github.com/sergiolaverde0/mangal/source"
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

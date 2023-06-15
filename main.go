package main

import (
	"github.com/metafates/mangal/cmd"
	"github.com/metafates/mangal/config"
	"github.com/metafates/mangal/log"
	"github.com/metafates/mangal/provider/generic/headless"
	"github.com/samber/lo"
)

func main() {
	transport := headless.GetTransportSingleton()
	defer func(transport *headless.Transport) {
		_ = transport.Close()
	}(transport)
	lo.Must0(config.Setup())
	lo.Must0(log.Setup())
	cmd.Execute()
}

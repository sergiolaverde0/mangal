package main

import (
	"github.com/metafates/mangal/cmd"
	"github.com/metafates/mangal/config"
	"github.com/metafates/mangal/log"
	"github.com/metafates/mangal/provider/generic/headless"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(config.Setup())
	lo.Must0(log.Setup())
	transport := headless.GetTransportSingleton()
	defer func(transport headless.TransportHeadless) {
		_ = transport.Close()
	}(transport)
	cmd.Execute()
}

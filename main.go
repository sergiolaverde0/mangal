package main

import (
	"github.com/belphemur/mangal/cmd"
	"github.com/belphemur/mangal/config"
	"github.com/belphemur/mangal/log"
	"github.com/belphemur/mangal/provider/generic/headless"
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

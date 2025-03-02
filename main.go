package main

import (
	"fmt"
	"github.com/sergiolaverde0/mangal/cmd"
	"github.com/sergiolaverde0/mangal/config"
	"github.com/sergiolaverde0/mangal/log"
	"github.com/sergiolaverde0/mangal/util"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(config.Setup())
	lo.Must0(log.Setup())
	defer util.Exit(0)

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		util.Exit(1)
	}
}

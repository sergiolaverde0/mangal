package main

import (
	"fmt"
	"github.com/belphemur/mangal/cmd"
	"github.com/belphemur/mangal/config"
	"github.com/belphemur/mangal/log"
	"github.com/belphemur/mangal/util"
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

package version

import (
	"fmt"
	"github.com/sergiolaverde0/mangal/color"
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/icon"
	"github.com/sergiolaverde0/mangal/key"
	"github.com/sergiolaverde0/mangal/style"
	"github.com/sergiolaverde0/mangal/util"
	"github.com/spf13/viper"
)

func Notify() {
	if !viper.GetBool(key.CliVersionCheck) {
		return
	}

	erase := util.PrintErasable(fmt.Sprintf("%s Checking if new version is available...", icon.Get(icon.Progress)))
	version, err := Latest()
	erase()
	if err == nil {
		comp, err := Compare(version, constant.Version)
		if err == nil && comp <= 0 {
			return
		}
	}

	fmt.Printf(`
%s New version is available %s %s
%s

`,
		style.Fg(color.Green)("▇▇▇"),
		style.Bold(version),
		style.Faint(fmt.Sprintf("(You're on %s)", constant.Version)),
		style.Faint("https://github.com/sergiolaverde0/mangal/releases/tag/v"+version),
	)

}

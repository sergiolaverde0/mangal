package cmd

import (
	"github.com/belphemur/mangal/color"
	"github.com/belphemur/mangal/config"
	"github.com/belphemur/mangal/constant"
	"github.com/belphemur/mangal/style"
	"github.com/belphemur/mangal/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.Flags().BoolP("set-only", "s", false, "only show variables that are set")
	envCmd.Flags().BoolP("unset-only", "u", false, "only show variables that are unset")

	envCmd.MarkFlagsMutuallyExclusive("set-only", "unset-only")
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Show available environment variables",
	Long:  `Show available environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		setOnly := lo.Must(cmd.Flags().GetBool("set-only"))
		unsetOnly := lo.Must(cmd.Flags().GetBool("unset-only"))

		config.EnvExposed = append(config.EnvExposed, where.EnvConfigPath)
		slices.Sort(config.EnvExposed)
		for _, env := range config.EnvExposed {
			if env != where.EnvConfigPath {
				env = strings.ToUpper(constant.Mangal + "_" + config.EnvKeyReplacer.Replace(env))
			}
			value := os.Getenv(env)
			present := value != ""

			if setOnly || unsetOnly {
				if !present && setOnly {
					continue
				}

				if present && unsetOnly {
					continue
				}
			}

			cmd.Print(style.New().Bold(true).Foreground(color.Purple).Render(env))
			cmd.Print("=")

			if present {
				cmd.Println(style.Fg(color.Green)(value))
			} else {
				cmd.Println(style.Fg(color.Red)("unset"))
			}
		}
	},
}

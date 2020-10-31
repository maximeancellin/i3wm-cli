package cmd

import (
	"github.com/maximeancellin/i3wm-cli/pkg/install"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tryCmd)
}

var tryCmd = &cobra.Command{
	Use:   "install",
	Short: "Install package to setup I3WM",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := installFunc(); err != nil {
			return err
		}
		return nil
	},
}

func installFunc() error {
	install.Packages("")
	return nil
}

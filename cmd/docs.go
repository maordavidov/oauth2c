package cmd

import (
	"github.com/cli/browser"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open docs page in browser",
	Run: func(cmd *cobra.Command, args []string) {
		_ = browser.OpenURL("https://github.com/maordavidov/oauth2c/blob/master/readme.md")
	},
}

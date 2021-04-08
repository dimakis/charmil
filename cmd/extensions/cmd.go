package extensions

import (
	"github.com/aerogear/charmil/cmd/extensions/install"
	"github.com/aerogear/charmil/cmd/extensions/installed"
	"github.com/aerogear/charmil/cmd/extensions/list"
	"github.com/aerogear/charmil/cmd/extensions/remove"
	"github.com/spf13/cobra"
)

// ExtensionsCmd represents the extensions command
var ExtensionsCmd = &cobra.Command{
	Use:   "extensions",
	Short: "A brief description of your command",
	Long:  ``,
}

func init() {
	ExtensionsCmd.AddCommand(install.InstallCmd, installed.InstalledCmd, list.ListCmd, remove.RemoveCmd)
}

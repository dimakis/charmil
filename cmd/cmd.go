package cmd

import (
	"github.com/spf13/cobra"

	"github.com/aerogear/charmil/cmd/extensions"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "charmil",
	Short: "A brief description of your application",
	Long:  ``,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(extensions.ExtensionsCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.charmil.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

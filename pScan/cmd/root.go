/*
Copyright © 2022 pqppq

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// used for flags
	cfgFile     string
	userLicence string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Version: "0.1",
		Use:     "pScan",
		Short:   "Fast TCP port scaner",
		Long:    `pScan - short for Port Scanner - executes TCP port scan on a list of hosts.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)

	// only verbose
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.pScan.yaml)")

	// specify short hand options
	rootCmd.PersistentFlags().StringP("hosts-file", "f", "pScan.hosts", "pScan hosts file")

}

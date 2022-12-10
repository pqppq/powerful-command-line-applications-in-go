/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "A brief description of your command",
	Long: `To load your completions run
source <(pScan  completion)
To load completions automatically on login, add this line to your .bashrc file:
$ ~/.bashrc
source <(pscan completion)
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return completionAction(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}

func completionAction(out io.Writer) error {
	return rootCmd.GenBashCompletion(out)
}

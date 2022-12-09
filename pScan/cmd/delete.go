/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/pqppq/pScan/scan"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:          "delete <host1>...<hostn>",
	Short:        "Delete host(s) from list",
	Aliases:      []string{"d"},
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return deleteAction(os.Stdout, hostFile, args)
	},
}

func init() {
	hostsCmd.AddCommand(deleteCmd)
}

func deleteAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}
	for _, h := range args {
		if err := hl.Remove(h); err != nil {
			return err
		}
		fmt.Fprintln(out, "Deleted host:", h)
	}
	return hl.Save(hostsFile)

}

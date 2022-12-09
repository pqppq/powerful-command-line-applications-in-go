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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:          "add <host1>...<hostn>",
	Short:        "Add new host(s) to list",
	Aliases:      []string{"a"},
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return addAction(os.Stdout, hostFile, args)
	},
}

func init() {
	hostsCmd.AddCommand(addCmd)
}

func addAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}
	for _, h := range args {
		if err := hl.Add(h); err != nil {
			return err
		}
		fmt.Fprintln(out, "Added host:", h)
	}
	return hl.Save(hostsFile)
}

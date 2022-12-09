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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List hosts in hosts list",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return listAction(os.Stdout, hostsFile, args)
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("list called")
	// },
	Aliases: []string{"l"},
}

func init() {
	hostsCmd.AddCommand(listCmd)
}

func listAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}
	for _, h := range hl.Hosts {
		if _, err := fmt.Fprintln(out, h); err != nil {
			return err
		}
	}
	return nil
}

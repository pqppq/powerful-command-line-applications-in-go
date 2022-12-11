/*
Copyright © 2022 pqppq

*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del <id>",
	Short: "Deletes an item from the list",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		return delAction(os.Stdout, apiRoot, args[0])
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func delAction(out io.Writer, apiRoot, arg string) error {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("%w: Item id must be a number", ErrNotNumber)
	}
	if err := deleteItem(apiRoot, id); err != nil {
		return err
	}
	return printDel(out, id)
}

func printDel(out io.Writer, id int) error {
	_, err := fmt.Fprintf(out, "Item number %d deleted\n", id)
	return err
}

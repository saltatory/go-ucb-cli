package cmd

import (
    "errors"
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(DeleteCommand)
}

var DeleteCommand = &cobra.Command{
    Use: "delete",
    Short: "Delete a new object",
    Long: "Delete a new object and return immediately",
    RunE: func(cmd *cobra.Command, args []string) error {
        return errors.New(ERROR_SUBCOMMAND)
    },
}


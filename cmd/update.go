package cmd

import (
    "errors"
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(UpdateCommand)
}

var UpdateCommand = &cobra.Command{
    Use: "update",
    Short: "Update an object",
    Long: "Update an object",
    RunE: func(cmd *cobra.Command, args []string) error {
         return errors.New(ERROR_SUBCOMMAND)
    },
}


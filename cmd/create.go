package cmd

import (
    "errors"
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(CreateCommand)
}

var CreateCommand = &cobra.Command{
    Use: "create",
    Short: "Create a new object",
    Long: "Create a new object and return immediately",
    RunE: func(cmd *cobra.Command, args []string) error {
        return errors.New(ERROR_SUBCOMMAND)
    },
}


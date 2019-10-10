package cmd

import (
    "errors"
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(GetCommand)
}

var GetCommand = &cobra.Command{
    Use: "get",
    Short: "Get an object",
    Long: "Get an object",
    RunE: func(cmd *cobra.Command, args []string) error {
         return errors.New(ERROR_SUBCOMMAND)
    },
}


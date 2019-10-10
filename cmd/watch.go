package cmd

import (
    "errors"
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(WatchCommand)
}

var WatchCommand = &cobra.Command{
    Use: "watch",
    Short: "Wait and watch",
    Long: "Wait on objects for state transitions",
    RunE: func(cmd *cobra.Command, args []string) error {
        return errors.New(ERROR_SUBCOMMAND)
    },
}


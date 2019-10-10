package cmd

import (
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(ListCommand)
}

var ListCommand = &cobra.Command{
    Use: "list",
    Short: "List objects",
    Long: "List objects",
}


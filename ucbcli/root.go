package ucbcli

import (
    "github.com/spf13/cobra"
)

var apiKey string

var rootCmd = &cobra.Command{
    Use:   "ucb",
    Short: "ucb makes Unity Cloud Build easy",
    Long:  "ucb is a CLI for Unity Cloud build",
    Run: func(cmd *cobra.Command, args []string) {
        //TODO
    },
}

func Execute() {
    // TODO
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&apiKey,"key","k","","Unity Cloud Build API Key")
}
package cmd

import (
    "github.com/spf13/cobra"
)

func init(){
    rootCmd.AddCommand(DownloadCommand)
}

var DownloadCommand = &cobra.Command{
    Use: "download",
    Short: "Download files",
    Long: "Download files",
}


package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func init() {
    ListCommand.AddCommand(ListBuildCommand)
}

var ListBuildCommand = &cobra.Command{
    Use:     "build",
    Short:   "List builds",
    Long:    "List builds",
    Aliases: []string{"builds"},
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildsApi

        b, _, err := builds.GetBuilds(ctx, OrgId, ProjectId, args[0], nil)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        for _,v := range b {
            fmt.Println(v.Build)
        }
    },
}

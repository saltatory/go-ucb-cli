package cmd

import (
    "context"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func init() {
    CreateCommand.AddCommand(CreateBuildCommand)
    CreateBuildCommand.Flags().StringVarP(&TargetId,"target","t","","Build Target Id (required)")
    CreateBuildCommand.MarkFlagRequired("target")
}

var ctx = context.TODO()

var CreateBuildCommand = &cobra.Command{
    Use:   "build",
    Short: "Create a new build",
    Long:  "Create a new build and return immediately",
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildsApi

        b, _, err := builds.StartBuilds(ctx, OrgId, ProjectId, TargetId, nil)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        for _,v := range b {
            fmt.Println(v.Build)
        }

    },
}

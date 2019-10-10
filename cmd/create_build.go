package cmd

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func init() {
    CreateCommand.AddCommand(CreateBuildCommand)
}

var ctx = context.TODO()

var CreateBuildCommand = &cobra.Command{
    Use:   "build",
    Short: "Create a new build",
    Long:  "Create a new build and return immediately",
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildsApi

        b, _, err := builds.StartBuilds(ctx, OrgId, ProjectId, "TODO", nil)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        fmt.Println(json.Marshal(b))

    },
}

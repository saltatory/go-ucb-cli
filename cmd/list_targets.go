package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func init() {
    ListCommand.AddCommand(GetTarget)
}

var GetTarget = &cobra.Command{
    Use:   "targets",
    Aliases: []string{"target"},
    Short: "Get all build targets",
    Long:  "Get all build targets",
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildtargetsApi

        b, _, err := builds.GetBuildTargetsForOrg(ctx, OrgId, nil)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        // TODO
        for _,v := range b {
            //fmt.Println(fmt.Sprintf("\"%s\",\"%s\"",v.Buildtargetid,v.Name))
            fmt.Println(v.Buildtargetid)
        }

    },
}

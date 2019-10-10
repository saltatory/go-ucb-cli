package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func init() {
    DeleteCommand.AddCommand(DeleteTarget)
}

var DeleteTarget = &cobra.Command{
    Use:   "target",
    Short: "Delete a build target",
    Long:  "Delete a build target",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildtargetsApi

        var errors = 0

        for _,v := range args {
            _, _, err := builds.OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidDelete(ctx, OrgId, ProjectId, v)

            if err != nil {
                errors++
                fmt.Println(err)
            } else {
                fmt.Println("Deleted Build Target Id " + v)
            }

            if (errors > 0) {
                os.Exit(1)
            }
        }


    },
}

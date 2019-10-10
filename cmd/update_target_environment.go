package cmd

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "strings"
)

var Environment string

const ENV_NAME string = "CI_ENVIRONMENT"

var ENV_VALID = map[string]string{"dev": "dev", "test": "test", "stage": "stage", "prod": "prod"}

func init() {
    UpdateTargetCommand.AddCommand(UpdateTargetEnvironment)
    UpdateTargetEnvironment.Flags().StringVarP(&Environment, "env", "e", "dev", "{dev|test|stage|prod}")
    UpdateTargetEnvironment.MarkFlagRequired("env")
}

var UpdateTargetEnvironment = &cobra.Command{
    Use:   "environment",
    Short: "Update a build target environment variables",
    Long:  "Update the environment variables for a build target",
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildtargetsApi

        Environment = strings.ToLower(Environment)

        _, ok := ENV_VALID[Environment]

        if !ok {
            fmt.Println("Invalid environment name")
            os.Exit(1)
        }

        env, _, err := builds.OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsGet(ctx, OrgId, ProjectId, TargetId)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        env[ENV_NAME] = Environment

        b, _, err := builds.OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsPut(ctx, OrgId, ProjectId, TargetId, env)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        fmt.Println(json.Marshal(b))

    },
}

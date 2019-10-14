package cmd

import (
    "fmt"
    ucbapi "github.com/saltatory/go-ucb-cli/api"
    "github.com/spf13/cobra"
    "os"
    "time"
)

const BUILD_STATUS_SUCCEEDED string = "success"
const BUILD_STATUS_FAILED string = "failure"
const BUILD_STATUS_PENDING string = "pending"
const BUILD_STATUS_QUEUED string = "queued"

const DEFAULT_TIMEOUT = 1800
const DEFAULT_POLL_INTERVAL = 5 * time.Second

var Timeout int

func init() {
    WatchCommand.AddCommand(WatchBuildCommand)
    WatchBuildCommand.Flags().IntVarP(&Timeout, "timeout", "T", DEFAULT_TIMEOUT, "Timeout in seconds")
    WatchBuildCommand.Flags().StringVarP(&TargetId, "target", "t", "", "Target Id (required)")
    _ = WatchBuildCommand.MarkFlagRequired("target")
}

var WatchBuildCommand = &cobra.Command{
    Use:   "build",
    Short: "Watch a build",
    Long:  "Watch a build until it completes",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildsApi

        c1 := make(chan string, 1)
        c2 := make(chan ucbapi.OrgsOrgidProjectsProjectidBuildtargetsBuilds, 1)

        go func() {
            timeout := time.Duration(Timeout) * time.Second
            time.Sleep(timeout)
            c1 <- "timeout"
        }()

        go func() {
            for {
                b, _, err := builds.GetBuild(ctx, OrgId, ProjectId, TargetId, args[0], nil)

                if err != nil {
                    fmt.Println(err)
                    os.Exit(1)
                }

                if b.BuildStatus != BUILD_STATUS_SUCCEEDED && b.BuildStatus != BUILD_STATUS_FAILED {
                    if b.BuildStatus != BUILD_STATUS_PENDING && b.BuildStatus != BUILD_STATUS_QUEUED{
                        now := time.Now().UTC()
                        started, err := time.Parse("2006-01-02T03:04:05.999Z", b.BuildStartTime)
                        if err != nil {
                            fmt.Println("Build time could not be comprehended")
                            os.Exit(1)
                        }
                        elapsed := now.UTC().Sub(started.UTC())
                        fmt.Print(fmt.Sprintf("\u001b[1000DDuration %s", elapsed.String()))
                    } else {
                        fmt.Println("Waiting for build to start.")
                    }
                    time.Sleep(DEFAULT_POLL_INTERVAL)
                    continue
                }

                c2 <- b
            }
        }()

        select {
        case <-c1:
            fmt.Println("Timeout exceeded")
            os.Exit(1)
        case build := <-c2:
            if build.BuildStatus == BUILD_STATUS_SUCCEEDED {
                fmt.Println(fmt.Sprintf("Build succeeded with %d errors and %d warnings.",build.BuildReport.Errors,build.BuildReport.Warnings))
                os.Exit(1)
            } else {
                fmt.Println(fmt.Sprintf("Build failed with %d errors and %d warnings.", build.BuildReport.Errors, build.BuildReport.Warnings))
                os.Exit(1)
            }
        }

    },
}

package cmd

import (
    "fmt"
    ucbapi "github.com/saltatory/go-ucb-cli/api"
    "github.com/spf13/cobra"
    "os"
    "time"
)

const BUILD_STATUS string = "buildStatus"
const BUILD_STATUS_SUCCEEDED string = "success"
const BUILD_STATUS_FAILED string = "failure"

const DEFAULT_TIMEOUT = 1800
const DEFAULT_POLL_INTERVAL = 5 * time.Second

var Timeout int

func init() {
    WatchCommand.AddCommand(WatchBuildCommand)
    WatchBuildCommand.Flags().IntVarP(&Timeout, "timeout", "t", DEFAULT_TIMEOUT, "Timeout in seconds")
}

var WatchBuildCommand = &cobra.Command{
    Use:   "build",
    Short: "Watch a build",
    Long:  "Watch a build until it completes",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildsApi

        c1 := make(chan string, 1)
        // TODO
        //c2 := make(chan map[string]interface{}, 1)
        c2 := make(chan ucbapi.OrgsOrgidProjectsProjectidBuildtargetsBuilds, 1)

        go func (){
            timeout := time.Duration(Timeout) * time.Second
            time.Sleep(timeout)
            c1 <- "timeout"
        }()

        go func() {
            for {
                b, _, err := builds.GetBuild(ctx, OrgId, ProjectId, TargetId, args[0], nil)

                if (err != nil) {
                    fmt.Println(err)
                    os.Exit(1)
                }

                //TODO
                if b.BuildStatus != BUILD_STATUS_SUCCEEDED && b.BuildStatus != BUILD_STATUS_FAILED {
                //if b[BUILD_STATUS] != BUILD_STATUS_SUCCEEDED && b[BUILD_STATUS] != BUILD_STATUS_FAILED {
                    time.Sleep(DEFAULT_POLL_INTERVAL)
                    continue;
                }

                c2 <- b;
            }
        }()

        select {
            case <-c1:
                fmt.Println("Timeout exceeded")
                os.Exit(1)
            case build :=<-c2:
                //TODO
                //fmt.Println(build["buildReport"])
                fmt.Println(build.BuildReport)
                //TODO
                //if build[BUILD_STATUS] == BUILD_STATUS_SUCCEEDED {
                if build.BuildStatus == BUILD_STATUS_SUCCEEDED {
                    fmt.Println("Build succeeded")
                    os.Exit(1)
                } else {
                    //TODO
                    //fmt.Println("Build failed. Status: ", build[BUILD_STATUS])
                    fmt.Println("Build failed. Status: ", build.BuildStatus)
                    os.Exit(1)
                }
        }


    },
}

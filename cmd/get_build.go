package cmd

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func init(){
    GetCommand.AddCommand(GetBuildCommand)
    GetBuildCommand.Flags().StringVarP(&TargetId,"target","t","","Build Target Id (required)")
    GetBuildCommand.MarkFlagRequired("target")
}

var GetBuildCommand = &cobra.Command{
    Use: "build",
    Short: "Get build information",
    Long: "Get status information about a build",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string){
        var builds = Client.BuildsApi

        b,_,err := builds.GetBuild(ctx,OrgId,ProjectId,TargetId,args[0],nil)

        if(err!=nil){
            fmt.Println(err)
            os.Exit(1)
        }

        val,_ := json.Marshal(b)
        fmt.Println(string(val))

    },
}


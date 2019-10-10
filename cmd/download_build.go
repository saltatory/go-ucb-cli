package cmd

import (
    "errors"
    "fmt"
    "github.com/spf13/cobra"
    "io"
    "net/http"
    "os"
    "path"
    "sync"
    "net/url"
)

func init() {
    DownloadCommand.AddCommand(DownloadBuildCommand)
    cwd,_ := os.Getwd()
    DownloadBuildCommand.Flags().StringVarP(&OutputDirectory, "dir", "d", cwd, "Output directory. Defaults to $PWD")
    DownloadBuildCommand.Flags().StringVarP(&TargetId,"target","t","","Build Target Id (required)")
    DownloadBuildCommand.MarkFlagRequired("target")
}

var DownloadBuildCommand = &cobra.Command{
    Use:   "build",
    Short: "Download artifacts from builds",
    Long:  "Download artifacts from builds",
    Run: func(cmd *cobra.Command, args []string) {
        var builds = Client.BuildsApi

        b,_,err := builds.GetBuild(ctx,OrgId,ProjectId,TargetId,args[0],nil)

        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }

        var waitgroup sync.WaitGroup

        for _,artifacts := range b.Links.Artifacts {
            for _,file := range artifacts.Files {
                waitgroup.Add(1)
                go DownloadFile(&waitgroup,OutputDirectory,file.Href)
            }
        }

        waitgroup.Wait()
    },
}

func DownloadFile(waitgroup *sync.WaitGroup, localpath string, remoteUrl string) error {
    defer waitgroup.Done()

    u,err :=  url.Parse(remoteUrl)
    filename := path.Join(localpath,path.Base(u.Path))

    fmt.Println("Starting download from '" + path.Base(u.Path) + "'")

    // Get the data
    resp, err := http.Get(remoteUrl)
    if err != nil {
        return err
    }
    if resp.StatusCode != 200 {
        return errors.New("Received failure code " + string(resp.StatusCode))
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filename)
    if err != nil {
        fmt.Println("Failed to download " + remoteUrl)
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}

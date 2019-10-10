package cmd

import (
    "errors"
    "fmt"
    ucbapi "github.com/saltatory/go-ucb-cli/api"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
)

const UCB_API_KEY string = "UCB_API_KEY"
const UCB_ORG_ID string = "UCB_ORG_ID"
const UCB_PROJECT_ID string = "UCB_PROJECT_ID"

var ApiKey string
var OrgId string
var ProjectId string
var TargetId string
var OutputDirectory string

var Config = ucbapi.NewConfiguration()
var Client = ucbapi.NewAPIClient(Config)
var Token string

const ERROR_SUBCOMMAND string = "Choose a sub-command"

var rootCmd = &cobra.Command{
    Use:   "ucb",
    Short: "ucb makes Unity Cloud Build easy",
    Long:  "ucb is a CLI for Unity Cloud build",
    RunE: func(cmd *cobra.Command, args []string) error {
        return errors.New(ERROR_SUBCOMMAND)
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        Config.DefaultHeader["Authorization"] = "Basic " + ApiKey
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}

func init() {
    viper.AutomaticEnv()
    rootCmd.PersistentFlags().StringVarP(&ApiKey, "key", "k", viper.GetString(UCB_API_KEY), "Unity Cloud Build API Key (required)")
    rootCmd.PersistentFlags().StringVarP(&OrgId, "org", "o", viper.GetString(UCB_ORG_ID), "Unity Cloud Build Org Id (required)")
    rootCmd.PersistentFlags().StringVarP(&ProjectId, "project", "p", viper.GetString(UCB_PROJECT_ID), "Unity Cloud Project Id (required)")
}

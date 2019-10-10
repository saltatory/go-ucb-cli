package cmd

import (
    "errors"
    "github.com/spf13/cobra"
)

func init(){
    UpdateCommand.AddCommand(UpdateTargetCommand)
    UpdateTargetCommand.PersistentFlags().StringVar(&TargetId,"id","","Build target id (required)")
    UpdateTargetCommand.MarkPersistentFlagRequired("id")
}

var UpdateTargetCommand = &cobra.Command{
    Use: "target",
    Short: "update a build target",
    Long: "Update a build target",
    RunE: func(cmd *cobra.Command, args []string) error {
         return errors.New("Choose a sub-command")
    },
}


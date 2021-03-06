package get

import (
	"github.com/jenkins-x/jx/v2/pkg/cmd/helper"
	"github.com/jenkins-x/jx/v2/pkg/cmd/opts"
	"github.com/jenkins-x/jx/v2/pkg/cmd/opts/step"
	get2 "github.com/jenkins-x/jx/v2/pkg/cmd/step/get"
	"github.com/spf13/cobra"
)

// StepGetOptions contains the command line flags
type StepGetOptions struct {
	step.StepOptions
}

// NewCmdStepGet Steps a command object for the "step" command
func NewCmdStepGet(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &StepGetOptions{
		StepOptions: step.StepOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get [command]",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	cmd.AddCommand(get2.NewCmdStepGetBuildNumber(commonOpts))
	cmd.AddCommand(get2.NewCmdStepGetVersionChangeSet(commonOpts))
	cmd.AddCommand(get2.NewCmdStepGetDependencyVersion(commonOpts))
	return cmd
}

// Run implements this command
func (o *StepGetOptions) Run() error {
	return o.Cmd.Help()
}

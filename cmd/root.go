package cmd

import "github.com/spf13/cobra"

// NewRootCmd creates the root command
func NewRootCmd() (c *cobra.Command) {
	c = &cobra.Command{
		Use:   "atest",
		Short: "API testing tool",
	}
	c.AddCommand(createInitCommand(),
		createRunCommand(), createSampleCmd())
	return
}
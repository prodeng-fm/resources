package cmd

import "github.com/spf13/cobra"

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage repository context (context.toml, classification, indexes)",
	Long: `The context command groups subcommands that operate on a
repository's context.toml — the file that declares the information
architecture rules harness enforces.

Subcommands will be added incrementally (init, validate, classify,
index, ...).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)
}

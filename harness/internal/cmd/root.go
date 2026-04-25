package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "harness",
	Short: "Harness — discipline for modern engineering repositories",
	Long: `harness is a tool for modern engineering harnessing:
information architecture, encoding discipline, policy as code,
controls, context lookup, and enterprise knowledge.

The initial focus is repository information architecture —
validating frontmatter, classifying documents, and maintaining
indexes — with the broader harness toolkit growing from there.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

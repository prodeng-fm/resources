package cmd

import (
	"bytes"
	"strings"
	"testing"
)

// execute runs the root command with the given args and returns combined
// stdout+stderr captured from cobra, plus any error from Execute.
func execute(t *testing.T, args ...string) (string, error) {
	t.Helper()
	buf := &bytes.Buffer{}
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return buf.String(), err
}

func TestRootAndContextPrintHelp(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"root no args", nil, "harness is a tool"},
		{"context no args", []string{"context"}, "groups subcommands that operate on a"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := execute(t, tc.args...)
			if err != nil {
				t.Fatalf("execute returned error: %v", err)
			}
			if !strings.Contains(out, tc.want) {
				t.Fatalf("output missing %q\n--- got ---\n%s", tc.want, out)
			}
		})
	}
}

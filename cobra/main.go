package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const ProgramName = "testcmd"

var (
	address    string
	tlsEnabled bool
)

func Cmd(programName string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   programName,
		Short: fmt.Sprintf("Sample %s program", programName),
		Long:  fmt.Sprintf("Long name of sample %s program", programName),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Welcome to %s world with address=[%s], tls=[%t]\n", programName, address, tlsEnabled)
			if len(args) > 0 {
				return fmt.Errorf("unexpected args are met: %v", args)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&address, "address", "d", "", "Target address")
	cmd.Flags().BoolVarP(&tlsEnabled, "tls", "", false, "Whethere TLS is enabled")

	return cmd
}

func main() {
	cmd := Cmd(ProgramName)
	if cmd.Execute() != nil {
		os.Exit(1)
	}
}

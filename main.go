package main

import (
	"fmt"
	"github.com/idoqo/sgh/command"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var expandedArgs []string
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}
	cmd, _, err := command.RootCmd.Traverse(expandedArgs)
	if err != nil || cmd == command.RootCmd {
		printError(err, cmd)
		os.Exit(2)
	}
	command.RootCmd.SetArgs(expandedArgs)
	if cmd, err := command.RootCmd.ExecuteC(); err != nil {
		 printError(err, cmd)
		 os.Exit(1)
	}
}

func printError(err error, cmd *cobra.Command) {
	fmt.Fprintf(os.Stderr, "failed to process command: %s\n", err)
}

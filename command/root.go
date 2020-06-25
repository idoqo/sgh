package command

import (
	"fmt"
	"github.com/idoqo/sgh/client"
	"github.com/spf13/cobra"
	"os"
)

var (
	SghVersion = "unknown"
	GoVersion  = "unknown"
	BuildDate  = "unknown"

	versionString = ""
)

// RootCmd is the cobra entry point and where the app starts executing
var RootCmd = &cobra.Command{
	Use: "sgh <command> <subcommand> [flags]",
	Short: "SwaggerHub CLI",
	Long: "Command line interface for the SwaggerHub API",

	Example: "$ sgh apis list",
}

var versionCmd = &cobra.Command{
	Use: "version",
	Hidden: true,
	Run: func(cmd *cobra.Command, args[]string) {
		fmt.Println(versionString)
	},
}

func init() {
	versionString = fmt.Sprintf("sgh %s (built: %s, go version: %s)", SghVersion, BuildDate, GoVersion)
	RootCmd.AddCommand(versionCmd)
}

func makeClient() (*client.Client, error) {
	token, err := readToken()
	if err != nil {
		return nil, err
	}

	return client.NewClient(token), nil
}

func readToken() (string, error){
	token := os.Getenv("SGH_TOKEN")
	if token == "" {
		err := fmt.Errorf("SGH_TOKEN environment variable not found")
		return "", err
	}
	return token, nil
}
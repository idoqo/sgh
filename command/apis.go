package command

import (
	"fmt"
	"github.com/idoqo/sgh/client"

	"github.com/spf13/cobra"
)

var apisCmd = &cobra.Command{
	Use:     "apis <subcommand> [flags]",
	Short:   "List, save and delete an API definition",
	Long:    "Work with SwaggerHub definitions",
	Example: "$ sgh apis list idoko/moira-oas",
	Annotations: map[string]string{
		"IsCore": "true",
		"help:arguments": `The API definition can be specified in any of the formats below:
- owner/name, e.g "idoko/moira-oas"`,
	},
}

func init() {
	RootCmd.AddCommand(apisCmd)
	apisCmd.AddCommand(defListCmd)
}

var defListCmd = &cobra.Command{
	Use: "list",
	Short: "List APIs created by a user",
	RunE: defList,
}

func defList(cmd *cobra.Command, args []string) error {
	apiClient, err := makeClient()
	if err != nil {
		return err
	}

	apis, err := client.GetAPIs(apiClient, "idoko")
	if err != nil {
		return err
	}

	for _, def := range apis.Apis {
		fmt.Printf("%s\n", def.Name)
	}
	return nil
}

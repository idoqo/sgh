package command

import "github.com/spf13/cobra"

var defCmd = &cobra.Command{
	Use: "def <command> [flags]",
	Short: "List, save and delete an API definition",
	Long: "Work with SwaggerHub definitions",
	Example: "$ sgh def list idoko/moira-oas",
	Annotations: map[string]string {
		"IsCore": "true",
		"help:arguments": `The API definition can be specified in any of the formats below:
- owner/name, e.g "idoko/moira-oas"`,
	},
}
func init() {
	RootCmd.AddCommand(defCmd)
}
